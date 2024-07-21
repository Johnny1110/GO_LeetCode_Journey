package taipei_mrt

//import (
//	"container/heap"
//	"fmt"
//	"math"
//)
//
//// Dijkstra function calculates the shortest path from start to target node.
//func Dijkstra_DEMO(start, target int) ([]int, int) {
//	dist := make(map[int]int)
//	prev := make(map[int]int)
//	pq := &PriorityQueue{}
//
//	heap.Init(pq)
//
//	for node := range graph {
//		dist[node] = math.MaxInt64
//		prev[node] = -1
//	}
//	dist[start] = 0
//	heap.Push(pq, Item{node: start, priority: 0})
//
//	for pq.Len() > 0 {
//		item := heap.Pop(pq).(Item)
//		u := item.node
//
//		if u == target {
//			break
//		}
//
//		for _, edge := range graph[u] {
//			v := edge.node
//			weight := edge.weight
//			alt := dist[u] + weight
//			if alt < dist[v] {
//				dist[v] = alt
//				prev[v] = u
//				heap.Push(pq, Item{node: v, priority: alt})
//			}
//		}
//	}
//
//	// Reconstruct the shortest path
//	path := []int{}
//	for u := target; u != -1; u = prev[u] {
//		path = append([]int{u}, path...)
//	}
//
//	if dist[target] == math.MaxInt64 {
//		return nil, -1
//	}
//	return path, dist[target]
//}
//
//func main() {
//	path, distance := Dijkstra(0, 4)
//	fmt.Printf("Shortest path from node 0 to node 4: %v with distance %d\n", path, distance)
//}
