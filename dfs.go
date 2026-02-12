package main

import (
	"errors"
)

func DepthFirstSearch(graph Graph, startNode, endNode int) []int {
	//graphSize := graph.Vertices
	visited := make([]bool, graph.Vertices)
	prnts := make([]int, 0)

	//for i := range graphSize {
	//	prnts[i] = NoParent
	//}

	if len(graph.GetNeighbors(startNode)) == 0 {
		panic(errors.New("starting vertex doesn't have any neighbors"))
	}

	path, _ := solveDFS(graph, startNode, endNode, visited, prnts)

	return path //reconstructPath(path, startNode, endNode)
}

func solveDFS(graph Graph, currNode, endNode int, visited []bool, path []int) ([]int, bool) {
	if visited[currNode] {
		return path, false
	}

	path = append(path, currNode)
	visited[currNode] = true

	if currNode == endNode {
		return path, true
	}

	for _, n := range graph.GetNeighbors(currNode) {
		var endNodeFound bool
		path, endNodeFound = solveDFS(graph, n, endNode, visited, path)
		if endNodeFound {
			return path, true
		}
	}

	return path, false
}
