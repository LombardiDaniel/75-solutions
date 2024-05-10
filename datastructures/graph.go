package datastructures

import "fmt"

type GraphAdjacencyListImpl struct {
	Vertices		[]*Vertex
}

type Vertex struct {
	Key 			int
	Adjacent		[]*Vertex
}

func NewGraph() GraphAdjacencyListImpl {
	return GraphAdjacencyListImpl{
		Vertices: []*Vertex{},
	} 
}

func (g *GraphAdjacencyListImpl) containsKey(key int) bool {
	for _, v := range g.Vertices {
		if v.Key == key {
			return true
		}
	}

	return false
}

func (g *GraphAdjacencyListImpl) AddVertex(key int) {

	if g.containsKey(key) {
		fmt.Println("CANNOT INSERT DUPLICATE KEY!!!")
		return
	}

	g.Vertices = append(g.Vertices, &Vertex{
		Key: key,
		Adjacent: []*Vertex{},
	})
}

func (g *GraphAdjacencyListImpl) AddEdge(from int, to int) {
	var vertexFrom *Vertex
	var vertexTo *Vertex
	for _, v := range g.Vertices {
		if v.Key == from {
			vertexFrom = v
		}
		if v.Key == to {
			vertexTo = v
		}
	}

	vertexFrom.Adjacent = append(vertexFrom.Adjacent, vertexTo)
}

func (g *GraphAdjacencyListImpl) Print() {
	for _, v := range g.Vertices {
		fmt.Printf("\nVertex %d : ", v.Key)
		for _, vv := range v.Adjacent {
			fmt.Printf("%d ", vv.Key)
		}
	}
	fmt.Println("")
}