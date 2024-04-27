package transform

import (
	"github.com/trevatk/anastasia/internal/adapter/port/rpc/msg"
	"github.com/trevatk/anastasia/internal/core/domain"
)

// Policy transform rpc policy into application policy
func Policy(in *msg.Policy) *domain.Policy {

	p := &domain.Policy{}

	p.Permission = domain.Permission(in.Permission)
	p.Signatures = in.Signatures

	var sbytes [28]byte
	copy(sbytes[:], []byte(in.Subject))
	p.Subject = sbytes

	return p
}
