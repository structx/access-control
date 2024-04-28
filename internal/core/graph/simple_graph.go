// Package graph implementation and provider
package graph

import (
	"fmt"
	"sync"

	"github.com/trevatk/anastasia/internal/core/domain"
	"golang.org/x/crypto/sha3"
)

// SimpleGraph dag implementation
type SimpleGraph struct {
	mtx      sync.RWMutex
	vertices map[[28]byte]domain.Vertex
	edges    map[[28]byte][]domain.Edge
}

// interface compliance
var _ domain.Graph = (*SimpleGraph)(nil)

// New constructor
func New() *SimpleGraph {
	return &SimpleGraph{
		mtx:      sync.RWMutex{},
		vertices: make(map[[28]byte]domain.Vertex),
		edges:    make(map[[28]byte][]domain.Edge),
	}
}

// AddVertex to graph
func (g *SimpleGraph) AddVertex(p *domain.Policy) (*domain.Vertex, error) {

	i := len(g.vertices)

	keyHash := hashKey([]byte(fmt.Sprintf("%d", i)))
	v := domain.Vertex{ID: keyHash, Tx: domain.Transaction{
		ID:        hashKey([]byte(fmt.Sprintf("%d", 0))),
		Subject:   [28]byte{},
		Resource:  keyHash,
		Signature: p.Signatures[0],
	}}

	g.vertices[keyHash] = v

	return &v, nil
}

// GetVertex from graph
func (g *SimpleGraph) GetVertex(key string) (*domain.Vertex, error) {

	keyHash := hashKey([]byte(key))
	v, ok := g.vertices[keyHash]
	if !ok {
		return nil, &ErrNotFound{Key: keyHash[:]}
	}

	return &v, nil
}

// AddEdge to vertex
func (g *SimpleGraph) AddEdge(source *domain.Vertex, target *domain.Vertex, p *domain.Policy) (*domain.Edge, error) {

	v, ok := g.vertices[target.ID]
	if !ok {
		return nil, &ErrNotFound{Key: source.Tx.ID[:]}
	}
	edges := g.edges[v.Tx.ID]

	e := domain.Edge{
		Subject:    p.Subject,
		Permission: p.Permission,
		Grantor:    p.Grantor,
		Signatures: p.Signatures,
	}

	// append edge policy to target vertex
	g.edges[v.Tx.ID] = append(edges, e)

	return &e, nil
}

// TraverseAndValidateData iterate over edges and validate data
func (g *SimpleGraph) TraverseAndValidateData(subject, resource string, permission domain.Permission) bool {

	keyHash := hashKey([]byte(resource))

	if _, ok := g.vertices[keyHash]; !ok {
		return false
	}

	es := g.edges[keyHash]
	if len(es) < 1 {
		// resource has no access policies
		// default policy deny
		return false
	}

	subj := hashKey([]byte(subject))

	for _, e := range es {
		if e.Subject == subj {
			if e.Permission == permission {
				return true
			}
		}
	}

	return false
}

func hashKey(key []byte) [28]byte {

	h := sha3.New224()
	h.Write(key)
	hash := h.Sum(nil)

	var result [28]byte
	copy(result[:], hash[:28])

	return result
}
