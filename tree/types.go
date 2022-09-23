package tree

type state string

const (
	VISITED state = "VISITED"
	VISTING state = "VISITING"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type path []int

type btree struct {
	root *Node
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
