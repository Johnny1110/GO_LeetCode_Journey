package taipei_mrt

import (
	"container/heap"
	"fmt"
	"math"
)

// Dijkstra function calculates the shortest path from start to target node.
func Dijkstra(start, target int) ([]int, int) {
	dist := make(map[int]int) // Distance from start to every single node
	prev := make(map[int]int) // Every single node previous one for shortest path
	pq := &PriorityQueue{}

	heap.Init(pq)

	// Init all the node data in dist & prev
	for node := range Graph {
		dist[node] = math.MaxInt32 // All node's distance from start set to infinity at first
		prev[node] = -1            // No Previous
	}

	// Init start node:
	dist[start] = 0                               // Distance from start to start is 0,
	heap.Push(pq, Item{node: start, priority: 0}) // And it will be the first choice

	for pq.Len() > 0 {
		item := heap.Pop(pq).(Item)
		currentNode := item.node

		if currentNode == target {
			break // the final
		}

		for _, edge := range Graph[currentNode] { // loop node currentNode's edge
			v := edge.node
			weight := edge.weight

			alt := dist[currentNode] + weight // alt is start ~ v node distance
			if alt < dist[v] {
				// if startNode -> someNode -> currentNode -> vNode
				// is better then
				// startNode -> someNode -> someNode -> vNode
				// (a better short cut is found):
				dist[v] = alt         // update new start ~ vNode distance
				prev[v] = currentNode // update new previous node to current
				heap.Push(pq, Item{node: v, priority: alt})
			}
		}
	}

	if dist[target] == math.MaxInt64 {
		// start node can't reach to the target node.
		return nil, -1
	}

	path := []int{}
	// find previous from target to start (reverse)
	for currentNode := target; currentNode != -1; currentNode = prev[currentNode] {
		path = append([]int{currentNode}, path...) // reverse append
		// path = append(path, currentNode)
	}

	return path, dist[target]
}

func Go() {
	path, dist := Dijkstra(0, 4)
	fmt.Println("Best way: ", path)
	fmt.Println("Distance from start to target: ", dist)
}
