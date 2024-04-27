package graph

import (
	"sync"

	"github.com/trevatk/anastasia/internal/core/domain"
)

// SimpleGraph
type SimpleGraph struct {
	mtx      sync.RWMutex
	vertices map[[28]byte]domain.Vertex
	edges    map[[28]byte][]domain.Edge
}

// New constructor
func New() *SimpleGraph {
	return &SimpleGraph{
		mtx:      sync.RWMutex{},
		vertices: make(map[[28]byte]domain.Vertex),
		edges:    make(map[[28]byte][]domain.Edge),
	}
}

// AddVertex
func (g *SimpleGraph) AddVertex() error {

	return nil
}

// AddEdge
func (g *SimpleGraph) AddEdge(source *domain.Vertex, target *domain.Vertex, p *domain.Policy) error {

	v, ok := g.vertices[target.ID]
	if !ok {
		return &ErrNotFound{Key: source.Tx.ID[:]}
	}
	edges := g.edges[v.Tx.ID]

	// append edge policy to target vertex
	g.edges[v.Tx.ID] = append(edges, domain.Edge{
		Subject:    p.Subject,
		Permission: p.Permission,
		Grantor:    p.Grantor,
		Signatures: p.Signatures,
	})

	return nil
}
