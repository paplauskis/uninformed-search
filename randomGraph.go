package main

import (
	"math/rand"
)

type RandomGraph struct {
	Vertices      int
	Edges         int
	AdjacencyList [][]int
}

const MAX_LIMIT = 30

func NewRandomGraph() *RandomGraph {

	vertices := 6 // rand.Intn(MAX_LIMIT) + 1
	maxEdges := computeMaxEdges(vertices)
	edges := rand.Intn(maxEdges) + 1

	adjList := make([][]int, vertices)
	for i := 0; i < vertices; i++ {
		adjList[i] = make([]int, 0)
	}

	graph := &RandomGraph{
		Vertices:      vertices,
		Edges:         edges,
		AdjacencyList: adjList,
	}

	// randomly generate edges (f.e. v is 2 and w is 5, so that means vertices 2 and 5 will be connected)
	for i := 0; i < edges; i++ {
		var v int
		var w int

		// prevents self-loops and duplicate edges
		for {
			v = rand.Intn(vertices)
			w = rand.Intn(vertices)

			if v != w && !graph.hasEdge(v, w) {
				break
			}
		}

		graph.addEdge(v, w)
	}

	return graph
}

func computeMaxEdges(numOfVertices int) int {
	return numOfVertices * (numOfVertices - 1) / 2
}

func (g *RandomGraph) addEdge(v, w int) {
	g.AdjacencyList[v] = append(g.AdjacencyList[v], w)
	g.AdjacencyList[w] = append(g.AdjacencyList[w], v)
}

func (g *RandomGraph) hasEdge(v, w int) bool {
	for _, x := range g.AdjacencyList[v] {
		if x == w {
			return true
		}
	}

	return false
}
