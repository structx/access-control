package domain

// Transaction
type Transaction struct {
	ID        [28]byte `json:"id"`
	Subject   [28]byte `json:"subject"`
	Resource  [28]byte `json:"resource"`
	Signature string   `json:"signature"`
}

// Vertex
type Vertex struct {
	ID [28]byte    `json:"id"`
	Tx Transaction `json:"tx"`
}

// Edge
type Edge struct {
	Subject    [28]byte   `json:"subject"`
	Permission Permission `json:"permission"`
	Grantor    [28]byte   `json:"grantor"`
	Signatures []string   `json:"signatures"`
}

// Graph
type Graph interface {
	AddVertex(*Policy) (*Vertex, error)
	GetVertex(resource string) (*Vertex, error)
	AddEdge(*Vertex, *Vertex, *Policy) (*Edge, error)
	TraverseAndValidateData(subject, resource string, permission Permission) bool
}
