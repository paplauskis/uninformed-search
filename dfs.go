package main

import (
	"errors"
)

func DepthFirstSearch(graph Graph, startNode, endNode int) []int {
	visited := make([]bool, graph.Vertices)
	history := make([]int, 0)

	if len(graph.GetNeighbors(startNode)) == 0 {
		panic(errors.New("starting vertex doesn't have any neighbors"))
	}

	path, _ := solveDFS(graph, startNode, endNode, visited, history)

	return path
}

func solveDFS(graph Graph, currNode, endNode int, visited []bool, history []int) ([]int, bool) {
	if visited[currNode] {
		return history, false
	}

	history = append(history, currNode)
	visited[currNode] = true

	if currNode == endNode {
		return history, true
	}

	for _, n := range graph.GetNeighbors(currNode) {
		var endNodeFound bool
		history, endNodeFound = solveDFS(graph, n, endNode, visited, history)
		if endNodeFound {
			return history, true
		}
	}

	return history, false
}
