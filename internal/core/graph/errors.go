package graph

import (
	"errors"
	"fmt"
)

var (
	// ErrMissingSignature source transaction signature not present in policy
	ErrMissingSignature = errors.New("source transaction signature missing from policy")
)

// ErrNotFound key not found
type ErrNotFound struct {
	Key []byte
}

// Error stringify error message
func (enf *ErrNotFound) Error() string {
	return fmt.Sprintf("key %x not found", enf.Key)
}
