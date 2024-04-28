package domain

// Transaction graph model
type Transaction struct {
	ID        [28]byte `json:"id"`
	Subject   [28]byte `json:"subject"`
	Resource  [28]byte `json:"resource"`
	Signature string   `json:"signature"`
}

// Vertex graph model
type Vertex struct {
	ID [28]byte    `json:"id"`
	Tx Transaction `json:"tx"`
}

// Edge graph model
type Edge struct {
	Subject    [28]byte   `json:"subject"`
	Permission Permission `json:"permission"`
	Grantor    [28]byte   `json:"grantor"`
	Signatures []string   `json:"signatures"`
}

// Graph functionality
type Graph interface {
	// AddVertex
	// TODO: change argument to transaction
	AddVertex(*Policy) (*Vertex, error)
	// GetVertex
	GetVertex(resource string) (*Vertex, error)
	// AddEdge
	// TODO:
	// change arguement from policy to transaction
	AddEdge(*Vertex, *Vertex, *Policy) (*Edge, error)
	// TraverseAndValidateData
	TraverseAndValidateData(subject, resource string, permission Permission) bool
}
