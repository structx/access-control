// Package service application logic
package service

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/trevatk/anastasia/internal/adapter/port/rpc/msg"
	"github.com/trevatk/anastasia/internal/core/domain"
	"github.com/trevatk/anastasia/internal/core/transform"
)

// AccessControl access controller implementation
type AccessControl struct {
	graph domain.Graph
}

// interface compliance
var _ domain.AccessController = (*AccessControl)(nil)

// NewAccessControl constructor
func NewAccessControl(simpleGraph domain.Graph) *AccessControl {
	return &AccessControl{
		graph: simpleGraph,
	}
}

// ModifyAccessControlList
func (ac *AccessControl) ModifyAccessControlList(ctx context.Context, modify *msg.ModifyACLPayload) (*msg.ModifyACLResponse, error) {

	// retrieve vertex of target
	target, err := ac.graph.GetVertex(modify.Resource)
	if err != nil {
		return nil, fmt.Errorf("failed to get vertex of resource %v", err)
	}

	// retrieve source of source
	source, err := ac.graph.GetVertex(modify.Subject)
	if err != nil {
		return nil, fmt.Errorf("failed to get vertex of subject %v", err)
	}

	// transform policy
	policy := transform.Policy(modify.Policy)

	// verify target signature is in policy
	ok := slices.Contains(policy.Signatures, target.Tx.Signature)
	if !ok {
		return nil, errors.New("policy does not include target signature")
	}

	// modify access control list
	_, err = ac.graph.AddEdge(source, target, policy)
	if err != nil {
		return nil, fmt.Errorf("failed to add edge %v", err)
	}

	return nil, nil
}
