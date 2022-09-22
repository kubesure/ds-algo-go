package tree

import (
	"container/list"
	"fmt"
)

func (g *graph) addVertex(v *Vertex) error {
	if g.vertices == nil {
		g.vertices = make([]*Vertex, 0)
	}
	g.vertices = append(g.vertices, v)
	return nil
}

func (g *graph) addAdjacent(v *Vertex, adjacent ...*Vertex) error {
	if v.adjacent == nil {
		v.adjacent = make([]*Vertex, 0)
	}
	v.adjacent = append(v.adjacent, adjacent...)
	return nil
}

func doDFS(n *Vertex, order *list.List) {

	for _, edge := range n.adjacent {
		if len(n.adjacent) > 0 {
			doDFS(edge, order)
		}
	}

	if n.state != VISITED {
		n.state = VISITED
		order.PushBack(n.key)
	}
}

func buildOrder(g *graph) []string {
	olist := &list.List{}
	for _, node := range g.vertices {
		doDFS(node, olist)
	}

	oarr := make([]string, 0)
	for {
		if olist.Len() > 0 {
			e := olist.Front()
			oarr = append(oarr, e.Value.(string))
			olist.Remove(e)
		} else {
			break
		}
	}
	return oarr
}

func NewGraph() *graph {
	return &graph{
		vertices: []*Vertex{},
	}
}

func (g *graph) print() {
	for _, v := range g.vertices {
		fmt.Printf("%v ", v.key)
		for _, a := range v.adjacent {
			fmt.Printf(" > %v ", a.key)
		}
		fmt.Println()
	}
}

func visit(n *Vertex) {
	fmt.Println(n.key)
}
