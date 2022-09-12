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
