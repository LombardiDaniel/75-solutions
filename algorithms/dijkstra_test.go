package algorithms_test

import (
	"fmt"
	"solutions/algorithms"
	"testing"
)

func TestDijkstra(t *testing.T) {
	// https://www.youtube.com/watch?v=_lHSawdgXpI
	graph := algorithms.NewWeightedGraph[string]()

	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")
	graph.AddVertex("E")

	graph.AddEdge("A", "B", 4)
	graph.AddEdge("A", "C", 2)
	graph.AddEdge("B", "C", 3)
	graph.AddEdge("C", "B", 1)
	graph.AddEdge("C", "D", 4)
	graph.AddEdge("B", "D", 2)
	graph.AddEdge("B", "E", 3)
	graph.AddEdge("C", "E", 5)
	graph.AddEdge("E", "D", 1)

	graph.Print()

	fmt.Printf("\n")
	fmt.Println(graph.Dijkstra("A", "C"))
	fmt.Println(graph.Dijkstra("A", "B"))
	fmt.Println(graph.Dijkstra("A", "D"))
	fmt.Println(graph.Dijkstra("A", "E"))
}