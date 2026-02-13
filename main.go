package main

import (
	"fmt"
	"time"
)

func main() {
	const iterations = 10000
	const graphC0unt = 10
	bfsElapsedArr := make([]time.Duration, graphC0unt)
	dfsElapsedArr := make([]time.Duration, graphC0unt)

	for i := 0; i < graphC0unt; i++ {
		graph := NewRandomGraph()

		start := time.Now()
		for i := 0; i < iterations; i++ {
			BreadthFirstSearch(*graph, 0, 5)
		}
		bfsElapsed := time.Since(start)
		bfsElapsedArr[i] = bfsElapsed

		start = time.Now()
		for i := 0; i < iterations; i++ {
			DepthFirstSearch(*graph, 0, 5)
		}
		dfsElapsed := time.Since(start)
		dfsElapsedArr[i] = dfsElapsed
	}

	fmt.Println("BFS elapsed array:", bfsElapsedArr)
	fmt.Println("DFS elapsed array:", dfsElapsedArr)

	//todo add analyser for results - avg, median...
}

func printGraph(graph Graph) {
	for i, list := range graph.AdjacencyList {
		fmt.Printf("%d -> { ", i)

		if len(list) == 0 {
			fmt.Print("No adjacent vertices")
		} else {
			for j, v := range list {
				fmt.Print(v)
				if j < len(list)-1 {
					fmt.Print(" , ")
				}
			}
		}

		fmt.Println(" }")
	}
}
