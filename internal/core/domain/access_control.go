// Package domain application models and interfaces
package domain

import (
	"context"
)

// Policy application model
type Policy struct {
	Subject    [28]byte
	Permission Permission
	Grantor    [28]byte
	Signatures []string
}

// AccessControlEntry
type AccessControlEntry struct {
	Subject, Resource string
	Permission        Permission
}

// NewServiceRegistered
type NewServiceRegistered struct {
	Policy *Policy
}

// ServiceRegistered
type ServiceRegistered struct{}

// UpdateRegisteredService
type UpdateRegisteredService struct{}

// UpdateAccessControlList
type UpdateAccessControlList struct {
	Resource string
	Subject  string
	Policy   *Policy
}

// AccessController required functionality for access service
type AccessController interface {
	// ServiceRegistered add new entry to access control list
	ServiceRegistered(context.Context, *NewServiceRegistered) (*ServiceRegistered, error)
	// ServiceUpdated todo
	ServiceUpdated(context.Context, *UpdateRegisteredService) error
	// ModifyAccessControlList update existing access control list entry
	ModifyAccessControlList(context.Context, *UpdateAccessControlList) error
	// VerifyServiceAccess
	VerifyServiceAccess(context.Context, *AccessControlEntry) error
	// VerifyUserAccess
	VerifyUserAccess(context.Context, *AccessControlEntry) error
}
