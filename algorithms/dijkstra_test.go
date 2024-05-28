package algorithms_test

import (
	"fmt"
	"solutions/algorithms"
	"testing"
)

func TestDijkstra(t *testing.T) {
	graph := algorithms.NewWeightedGraph[string]()

	err := graph.AddVertex("A")
	if err != nil {
		t.Errorf(err.Error())
	}
	err = graph.AddVertex("B")
	if err != nil {
		t.Errorf(err.Error())
	}
	err = graph.AddVertex("C")
	if err != nil {
		t.Errorf(err.Error())
	}
	graph.AddVertex("D")
	graph.AddVertex("E")

	err = graph.AddEdge("A", "B", 4)
	if err != nil {
		t.Errorf(err.Error())
	}
	err = graph.AddEdge("A", "C", 2)
	if err != nil {
		t.Errorf(err.Error())
	}
	err = graph.AddEdge("B", "C", 3)
	if err != nil {
		t.Errorf(err.Error())
	}
	err = graph.AddEdge("C", "B", 1)
	if err != nil {
		t.Errorf(err.Error())
	}

	graph.AddEdge("C", "D", 4)
	graph.AddEdge("B", "D", 2)
	graph.AddEdge("B", "E", 3)
	graph.AddEdge("C", "E", 5)
	graph.AddEdge("E", "D", 1)

	graph.Print()

	fmt.Println(graph.Dijkstra("A", "C"))
	fmt.Println(graph.Dijkstra("A", "B"))
	fmt.Println(graph.Dijkstra("A", "D"))
}