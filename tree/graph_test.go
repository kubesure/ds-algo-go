package tree

import (
	"testing"
)

func TestBuildOrder1(t *testing.T) {
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

	if len(o) != 7 {
		t.Errorf("order should have 7 tasks")
	}

}

func TestBuildOrder2(t *testing.T) {
	g := NewGraph()

	a := &Vertex{key: "5"}
	b := &Vertex{key: "4"}
	c := &Vertex{key: "3"}
	d := &Vertex{key: "2"}
	e := &Vertex{key: "1"}
	f := &Vertex{key: "0"}

	g.addVertex(a)
	g.addVertex(b)
	g.addVertex(c)
	g.addVertex(d)
	g.addVertex(e)
	g.addVertex(f)

	g.addAdjacent(a, d, f)
	g.addAdjacent(b, f, e)
	g.addAdjacent(d, c)
	g.addAdjacent(c, e)

	g.print()

	o := buildOrder(g)

	if len(o) != 6 {
		t.Errorf("order should have 6 tasks")
	}

}
