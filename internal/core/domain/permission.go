package domain

// Permission
type Permission string

const (
	// Read
	Read Permission = "read"
	// Write
	Write Permission = "write"
	// Delete
	Delete Permission = "delete"
	// Execute
	Execute Permission = "execute"
)

// String
func (p Permission) String() string {
	return string(p)
}
