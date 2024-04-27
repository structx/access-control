// Package domain application models and interfaces
package domain

import (
	"context"

	"github.com/trevatk/anastasia/internal/adapter/port/rpc/msg"
)

// Policy application model
type Policy struct {
	Subject    [28]byte
	Permission Permission
	Grantor    [28]byte
	Signatures []string
}

// AccessController required functionality for access service
type AccessController interface {
	// ModifyAccessControlList create a new policy and attach to target edges
	ModifyAccessControlList(ctx context.Context, msg *msg.ModifyACLPayload) (*msg.ModifyACLResponse, error)
}
