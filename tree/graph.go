package tree

import "fmt"

func (g *graph) addNode(n *GNode) error {
	if g.nodes == nil {
		g.nodes = make([]*GNode, 0)
	}
	g.nodes = append(g.nodes, n)
	return nil
}

func (g *graph) addEdges(n *GNode, edge ...*GNode) error {
	if g.edges == nil {
		g.edges = make(map[*GNode][]*GNode)
	}
	g.edges[n] = append(g.edges[n], edge...)
	return nil
}

func (g *graph) print() {
	for node, edges := range g.edges {
		fmt.Printf("%v ", node.data)
		for _, v := range edges {
			fmt.Printf(" > %v ", v.data)
		}
		fmt.Println()
	}
}

func NewGraph() *graph {
	return &graph{}
}
