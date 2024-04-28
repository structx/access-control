// Package service application logic
package service

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/trevatk/anastasia/internal/core/domain"
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

// ModifyAccessControlList update existing access control list entry
func (ac *AccessControl) ModifyAccessControlList(_ context.Context, op *domain.UpdateAccessControlList) error {

	// retrieve vertex of target
	target, err := ac.graph.GetVertex(op.Resource)
	if err != nil {
		return fmt.Errorf("failed to get vertex of resource %v", err)
	}

	// retrieve source of source
	source, err := ac.graph.GetVertex(op.Subject)
	if err != nil {
		return fmt.Errorf("failed to get vertex of subject %v", err)
	}

	// verify target signature is in policy
	ok := slices.Contains(op.Policy.Signatures, target.Tx.Signature)
	if !ok {
		return errors.New("policy does not include target signature")
	}

	// modify access control list
	_, err = ac.graph.AddEdge(source, target, op.Policy)
	if err != nil {
		return fmt.Errorf("failed to add edge %v", err)
	}

	return nil
}

// ServiceRegistered add new entry to access control list
func (ac *AccessControl) ServiceRegistered(_ context.Context, op *domain.NewServiceRegistered) (*domain.ServiceRegistered, error) {

	_, err := ac.graph.AddVertex(op.Policy)
	if err != nil {
		return nil, fmt.Errorf("unable to add vertex %v", err)
	}

	return &domain.ServiceRegistered{}, nil
}

// ServiceUpdated todo
func (ac *AccessControl) ServiceUpdated(_ context.Context, _ *domain.UpdateRegisteredService) error {
	// TODO:
	// implement handler
	return nil
}

// VerifyServiceAccess with graph
func (ac *AccessControl) VerifyServiceAccess(_ context.Context, ace *domain.AccessControlEntry) error {

	granted := ac.graph.TraverseAndValidateData(ace.Subject, ace.Resource, ace.Permission)
	if granted {
		return nil
	}

	return errors.New("access denied")
}

// VerifyUserAccess with graph
func (ac *AccessControl) VerifyUserAccess(_ context.Context, ace *domain.AccessControlEntry) error {

	granted := ac.graph.TraverseAndValidateData(ace.Subject, ace.Resource, ace.Permission)
	if granted {
		return nil
	}

	return errors.New("access denied")
}
