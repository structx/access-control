package msg

import "time"

// Policy controller model
type Policy struct {
	Subject    string     `json:"subject"`
	Permission string     `json:"permission"`
	Grantor    string     `json:"grantor"`
	Signatures []string   `json:"signatures"`
	ExpiresAt  *time.Time `json:"expires_at,omitempty"`
}
