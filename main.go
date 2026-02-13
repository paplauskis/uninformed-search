package main

import (
	"fmt"
	"slices"
	"time"
)

func main() {
	const iterations = 10000
	const graphCount = 100
	bfsElapsedArr := make([]int64, graphCount)
	dfsElapsedArr := make([]int64, graphCount)

	for i := 0; i < graphCount; i++ {
		graph := NewRandomGraph()

		start := time.Now()
		for i := 0; i < iterations; i++ {
			BreadthFirstSearch(*graph, 0, 5)
		}
		bfsElapsed := time.Since(start)
		bfsElapsedArr[i] = bfsElapsed.Milliseconds()

		start = time.Now()
		for i := 0; i < iterations; i++ {
			DepthFirstSearch(*graph, 0, 5)
		}
		dfsElapsed := time.Since(start)
		dfsElapsedArr[i] = dfsElapsed.Milliseconds()
	}

	fmt.Println("BFS elapsed array:", bfsElapsedArr)
	fmt.Println("DFS elapsed array:", dfsElapsedArr)

	bfsStats := &Stats{
		Average: calcAverage(bfsElapsedArr),
		Median:  calcMedian(bfsElapsedArr),
		Min:     slices.Min(bfsElapsedArr),
		Max:     slices.Max(bfsElapsedArr),
	}

	dfsStats := &Stats{
		Average: calcAverage(dfsElapsedArr),
		Median:  calcMedian(dfsElapsedArr),
		Min:     slices.Min(dfsElapsedArr),
		Max:     slices.Max(dfsElapsedArr),
	}

	fmt.Printf("BFS times in ms: avg - %f, median - %f, min - %d, max - %d\n", bfsStats.Average, bfsStats.Median, bfsStats.Min, bfsStats.Max)
	fmt.Printf("DFS times in ms: avg - %f, median - %f, min - %d, max - %d\n", dfsStats.Average, dfsStats.Median, dfsStats.Min, dfsStats.Max)
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
