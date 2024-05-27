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

func containsKey(verices []*Vertex, key int) bool {
	for _, v := range verices {
		if v.Key == key {
			return true
		}
	}

	return false
}

func (g *GraphAdjacencyListImpl) GetVertex(key int) *Vertex {
	for _, v := range g.Vertices {
		if v.Key == key {
			return v
		}
	}

	return nil
}

func (g *GraphAdjacencyListImpl) AddVertex(key int) {

	if containsKey(g.Vertices, key) {
		fmt.Println("CANNOT INSERT DUPLICATE KEY!!!")
		return
	}

	g.Vertices = append(g.Vertices, &Vertex{
		Key: key,
		Adjacent: []*Vertex{},
	})
}

func (g *GraphAdjacencyListImpl) AddEdge(from int, to int) {
	vertexFrom := g.GetVertex(from)
	vertexTo := g.GetVertex(to)

	if vertexFrom == nil || vertexTo == nil {
		fmt.Printf("Invalid edge (%d -> %d)\n", from, to)
		return
	}

	if containsKey(vertexFrom.Adjacent, to) {
		fmt.Printf("Edge (%d -> %d) already exists\n", from, to)
		return
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

func (g *GraphAdjacencyListImpl) DepthFirstSearch(val int) *Vertex {
	
	visited := make(map[*Vertex]bool)
	for _, v := range g.Vertices {
		vertex := exploreVertex(v, val, &visited)
		if vertex != nil {
			return vertex
		}
	}
	return nil
}

func exploreVertex(vertex *Vertex, val int, visited *map[*Vertex]bool) *Vertex {
	if vertex.Key == val {
		return vertex
	}

	for _, v := range vertex.Adjacent {
		(*visited)[v] = true
		return exploreVertex(v, val, visited)
	}

	return nil
}

func (g *GraphAdjacencyListImpl) NumPaths(from int, to int) int {

	var start *Vertex
	var end *Vertex
	for _, v := range g.Vertices {
		if v.Key == from {
			start = v
		}
		if v.Key == to {
			end = v
		}
	}

	if start == nil || end == nil {
		return 0
	}

	visited := make(map[int]bool)
	return numPaths(start, end, visited)
}

func numPaths(curr *Vertex, tgt *Vertex, visitedKeys map[int]bool) int {
	if curr == tgt {
		return 1
	}

	visitedKeys[curr.Key] = true
	ret := 0
	for _, v := range curr.Adjacent {
		visited, exists := visitedKeys[v.Key]
		if !exists || !visited {
			ret += numPaths(v, tgt, visitedKeys)
		}
	}
	visitedKeys[curr.Key] = false

	return ret
}

func LargestPath(arr []int) int {

	LISs := []int{}
	LISs = append(LISs, arr...)

	for i := len(arr) - 1; i >= 0; i-- {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				LISs[i] = max(LISs[i], 1 + LISs[j])
			}
		}
	}

	maxVal := 1
	for _, v := range LISs {
		if v > maxVal {
			maxVal = v
		}
	}

	return maxVal
}