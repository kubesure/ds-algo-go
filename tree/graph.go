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
		v.adjacent = make(map[string]*Vertex, 0)
	}
	for _, a := range adjacent {
		if _, ok := v.adjacent[a.key]; !ok {
			v.adjacent[a.key] = a
		} else {
			return fmt.Errorf("Vertex %v existing in graph", a.key)
		}
	}

	return nil
}

func doDFS(v *Vertex, order *list.List) {

	if v.state == VISITED {
		return
	}

	for _, a := range v.adjacent {
		if len(v.adjacent) > 0 {
			doDFS(a, order)
		}
	}

	if v.state != VISITED {
		v.state = VISITED
		order.PushBack(v.key)
		visit(v)
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
	fmt.Printf("%v ", n.key)
}
