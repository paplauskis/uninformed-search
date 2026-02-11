package main

import "fmt"

func main() {
	graph := NewRandomGraph()

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

	vertices := BreadthFirstSearch(*graph, 0, 5)
	fmt.Println(vertices)
}
