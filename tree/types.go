package tree

type state string

const (
	VISITED state = "VISITED"
	VISTING state = "VISITING"
)

type Node struct {
	value, index, size  int
	parent, left, right *Node
}

type path []int
type elements []int

type btree struct {
	root *Node
	size int32
}

type Vertex struct {
	key string
	//adjacent []*Vertex
	adjacent map[string]*Vertex
	state    state
}

type graph struct {
	vertices []*Vertex
}
