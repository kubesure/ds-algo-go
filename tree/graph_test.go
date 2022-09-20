package tree

import "testing"

func TestBuildGraph(t *testing.T) {
	g := NewGraph()

	a := &GNode{data: "A"}
	b := &GNode{data: "B"}
	c := &GNode{data: "C"}
	d := &GNode{data: "D"}
	e := &GNode{data: "E"}
	f := &GNode{data: "F"}
	h := &GNode{data: "H"}

	g.addNode(a)
	g.addNode(b)
	g.addNode(c)
	g.addNode(d)
	g.addNode(e)
	g.addNode(f)
	g.addNode(h)

	g.addEdges(a, c)
	g.addEdges(c, d, f)

	g.print()

}
