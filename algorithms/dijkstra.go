package algorithms

import (
	"errors"
	"fmt"
	"math"
	"slices"
)


type weightedGraphVertex[T comparable] struct {
	edges			[]weightedGraphEdge[T]
	key				T
}

type weightedGraphEdge[T comparable] struct {
	cost			float32
	vertex			*weightedGraphVertex[T]
}

type weightedGraphRoute[T comparable] struct {
	cost			float32
	vertex			*weightedGraphVertex[T]
	prev			*weightedGraphVertex[T]
}

type weightedGraphRouteRet[T comparable] struct {
	cost			float32
	vertexKey			T
}

type WeightedGraph[T comparable] struct {
	vertices []weightedGraphVertex[T]
}

func NewWeightedGraph[T comparable]() WeightedGraph[T] {
	return WeightedGraph[T]{
		vertices: []weightedGraphVertex[T]{},
	}
}

func (g *WeightedGraph[T]) AddVertex(key T) error {
	_, err := g.GetVertex(key)
	if err == nil {
		return errors.New("key already in graph")
	}

	newVertex := weightedGraphVertex[T]{
		edges: []weightedGraphEdge[T]{},
    	key: key,
	}

	g.vertices = append(g.vertices, newVertex)

	return nil
}

func (g *WeightedGraph[T]) GetVertex(key T) (*weightedGraphVertex[T], error) {
	for _, v := range g.vertices {		
		if v.key == key {
			return &v, nil
		}
	}

	return nil, errors.New("key already in graph")
}

func (g *WeightedGraph[T]) AddEdge(from T, to T, cost float32) error {
	toVertex, err := g.GetVertex(to)
	if err != nil {
		return errors.New("'to' already in graph")
	}

	for i, v := range g.vertices {
		if v.key == from {
			g.vertices[i].edges = append(v.edges, weightedGraphEdge[T]{
				cost: cost,
				vertex: toVertex,
			})
			return nil
		}
	}

	return errors.New("'from' not in graph")
}

func (g *WeightedGraph[T]) Print() {
	fmt.Print("Vertices::")
	for _, v := range g.vertices {
		fmt.Printf("\n%v :", v.key)
		for _, vv := range v.edges {
			fmt.Printf(" %.2f -> %v |", vv.cost, vv.vertex.key)
		}
	}
	fmt.Print("\n")
}

func (g *WeightedGraph[T]) Dijkstra(from T, to T) ([]weightedGraphRouteRet[T], error) {

	path := make(map[T]*weightedGraphRoute[T])

	costs := make(map[T]float32)
	parents := make(map[T]T)
	unvisiteds := make(map[T]bool)

	fromVertex, err := g.GetVertex(from)
	if err != nil {
		return nil, err
	}

	// set initial costs to infinte & unvisiteds
	for _, v := range g.vertices {
		costs[v.key] = float32(math.MaxFloat32)
		unvisiteds[v.key] = true
	}

	costs[from] = 0
	path[from] = &weightedGraphRoute[T]{
		cost: 0,
		vertex: fromVertex,
		prev: nil,
	}
	currVertex := fromVertex
	for currVertex != nil {
		// fmt.Printf("currVertex: %v\n", currVertex)
		currCost := costs[currVertex.key]
		for _, v := range currVertex.edges {
			newCost := currCost + v.cost
			// fmt.Printf("%v newCost %v: %v + %v = %v | old = %v | update?=%v\n", currVertex.key, v.vertex.key, currCost, v.cost, newCost, costs[v.vertex.key], newCost < costs[v.vertex.key])
			if newCost < costs[v.vertex.key] {
				costs[v.vertex.key] = newCost
				parents[v.vertex.key] = currVertex.key
				path[v.vertex.key] = &weightedGraphRoute[T]{
					cost: newCost,
					vertex: v.vertex,
					prev: currVertex,
				}
			}
		}

		unvisiteds[currVertex.key] = false
		currVertex = g.getLowestCostNeighbour(currVertex, unvisiteds)
	}

	// fmt.Printf("costs: %v\n", costs)

	ret := []weightedGraphRouteRet[T]{}
	currPath := path[to]
	for currPath != nil {
		// fmt.Printf("currPath: %v\n", currPath)
		ret = append(ret, weightedGraphRouteRet[T]{
			cost: currPath.cost,
    		vertexKey: currPath.vertex.key,
		})
		if currPath.prev == nil {
			currPath = nil
		} else {
			currPath = path[currPath.prev.key]
		}
	}
	slices.Reverse(ret)
	return ret, nil
}

func (g *WeightedGraph[T]) getLowestCostNeighbour(vertex *weightedGraphVertex[T], unvisiteds map[T]bool) *weightedGraphVertex[T] {
	lowestCost := float32(math.MaxFloat32)
	var lowestCostVertex *weightedGraphVertex[T] = nil
	for _, v := range vertex.edges {
		if v.cost < lowestCost && unvisiteds[v.vertex.key] {
			lowestCost = v.cost
			lowestCostVertex = v.vertex
		}
	}

	if lowestCostVertex == nil {
		return nil
	}

	// WTF?::?
	v, err := g.GetVertex(lowestCostVertex.key)
	if err != nil {
		return nil
	}

	return v
}