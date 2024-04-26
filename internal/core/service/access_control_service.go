// Package service application logic
package service

import "github.com/trevatk/anastasia/internal/core/domain"

// AccessControl access controller implementation
type AccessControl struct{}

// interface compliance
var _ domain.AccessController = (*AccessControl)(nil)

// NewAccessControl constructor
func NewAccessControl() *AccessControl {
	return &AccessControl{}
}
