package find_median_from_data_stream

import "fmt"

type Queue struct {
	nodes []*Node
}

func (q *Queue) Push(n *Node) {
	q.nodes = append(q.nodes, n)
}

func (q *Queue) Pop() *Node {
	if len(q.nodes) == 0 {
		return nil
	}
	n := q.nodes[0]
	q.nodes = q.nodes[1:]
	return n
}

func (q *Queue) Empty() bool {
	return len(q.nodes) == 0
}

func (q *Queue) Size() int {
	return len(q.nodes)
}

func (n *Node) Debug() {
	q := &Queue{
		nodes: make([]*Node, 0),
	}

	q.Push(n)

	currentLevel := 0
	for !q.Empty() {
		currentLevelLen := q.Size()
		fmt.Printf("Level - [%v] ==========================\n", currentLevel)
		for range currentLevelLen {
			node := q.Pop()

			if node != nil {
				fmt.Printf("[%v] ", node.Val)
			} else {
				fmt.Printf("[ ] ")
			}

			q.Push(node.Left)
			q.Push(node.Right)
		}
		fmt.Printf("\n")
		currentLevel++
	}
}
