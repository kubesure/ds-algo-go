package tree

type state string

const (
	VISITED state = "VISITED"
	VISTING state = "VISITING"
)

type Node struct {
	value       int
	left, right *Node
	nodeState   state
}
