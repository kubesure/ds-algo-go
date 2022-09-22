package tree

import (
	"fmt"
	"testing"
)

func TestBuildGraph(t *testing.T) {
	g := NewGraph()

	a := &Vertex{key: "A"}
	b := &Vertex{key: "B"}
	c := &Vertex{key: "C"}
	d := &Vertex{key: "D"}
	e := &Vertex{key: "E"}
	f := &Vertex{key: "F"}
	h := &Vertex{key: "H"}

	g.addVertex(e)
	g.addVertex(f)
	g.addVertex(h)
	g.addVertex(a)
	g.addVertex(b)
	g.addVertex(c)
	g.addVertex(d)

	g.addAdjacent(a, c, h)
	g.addAdjacent(c, d, f)

	g.print()

	o := buildOrder(g)

	for _, v := range o {
		fmt.Printf("%v ", v)
	}

	if len(o) != 7 {
		t.Errorf("order should have 7 tasks")
	}

}
