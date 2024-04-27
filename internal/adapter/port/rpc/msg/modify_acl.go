package msg

import (
	"errors"
	"time"
)

// ModifyACLPayload
type ModifyACLPayload struct {
	Subject  string  `json:"subject"`
	Resource string  `json:"resource"`
	Policy   *Policy `json:"policy"`
}

// ModifyACLResponse
type ModifyACLResponse struct{}

// Validate
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
