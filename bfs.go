package main

import "errors"

const NoParent = -1

func BreadthFirstSearch(graph Graph, startNode int, endNode int) ([]int, []int) {
	if startNode < 0 ||
		endNode < 0 {
		return nil, nil
	}

	if len(graph.GetNeighbors(startNode)) == 0 {
		panic(errors.New("starting vertex doesn't have any neighbors"))
	}

	path, history := solveBFS(graph, startNode, endNode)

	return reconstructPath(path, startNode, endNode), history
}

func solveBFS(graph Graph, startNode, endNode int) ([]int, []int) {
	graphSize := graph.Vertices
	visited := make([]bool, graphSize)

	//stores parents, f.e. if vertex 0 is the starting point and it goes to vertex 3,
	//vertex 3's parent is 0, if vertex 3 has a child vertex 5 (for example),
	//then vertex 5's parent is 3 and so on....
	prnts := make([]int, graphSize)
	history := []int{0}

	for i := range graphSize {
		prnts[i] = NoParent
	}

	queue := NewQueue()
	visited[startNode] = true
	queue.Enqueue(startNode)

	for {
		if queue.Size == 0 {
			return prnts, history
		}
		node, _ := queue.Dequeue()
		neighbors := graph.AdjacencyList[node]

		for _, x := range neighbors {
			if visited[x] {
				continue
			}

			queue.Enqueue(x)
			visited[x] = true
			prnts[x] = node
			history = append(history, x)

			if x == endNode {
				return prnts, history
			}
		}
	}
}

// reconstructs path from vertex's parents (latest search)
//
//	0  1  2  3  4  5  (array indexes)
//
// [-1, 0, 0, 1, 0, 3] (array values, -1 is root vertex)
// if path needs to be found from vertex 0 to vertex 3, then returned array
// will be [0, 1, 3] as a graph it will look like 0 -> 1 -> 3
func reconstructPath(prnts []int, startNode, endNode int) []int {
	var path []int
	for i := endNode; i != NoParent; i = prnts[i] {
		path = append(path, i)
	}

	reversedPath := reverse(path) //reverse for readability
	if reversedPath[0] == startNode {
		return reversedPath
	}

	return []int{}
}

func reverse(arr []int) []int {
	newArr := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		newArr[i] = arr[len(arr)-1-i]
	}

	return newArr
}
