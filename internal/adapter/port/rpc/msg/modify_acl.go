// Package msg controller models
package msg

import (
	"errors"
	"time"
)

// ModifyACLPayload controller model
type ModifyACLPayload struct {
	Subject  string  `json:"subject"`
	Resource string  `json:"resource"`
	Policy   *Policy `json:"policy"`
}

// ModifyACLResponse controller model
type ModifyACLResponse struct{}

// Validate payload
func (macl *ModifyACLPayload) Validate() error {

	if macl.Policy == nil {
		return errors.New("empty policy")
	}

	if macl.Policy.ExpiresAt != nil {
		if macl.Policy.ExpiresAt.After(time.Now()) {
			return errors.New("invalid timestamp")
		}
	}

	return nil
}
