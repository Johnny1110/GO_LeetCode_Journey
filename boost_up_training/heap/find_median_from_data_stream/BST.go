package find_median_from_data_stream

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Size  int // total nodes in this subtree (including self) 1 + leftSubSize + rightSubSize
}

func NewNode(val int) *Node {
	return &Node{
		Val:  val,
		Size: 1,
	}
}

func (n *Node) getLeftSize() int {
	if n.Left == nil {
		return 0
	}
	return n.Left.Size
}

func (n *Node) getRightSize() int {
	if n.Right == nil {
		return 0
	}
	return n.Right.Size
}

func (n *Node) Push(val int) {
	n.Size += 1
	currentVal := n.Val
	if currentVal >= val { // eq also goes left
		if n.Left != nil {
			n.Left.Push(val)
		} else {
			n.Left = NewNode(val)
		}
	} else {
		if n.Right != nil {
			n.Right.Push(val)
		} else {
			n.Right = NewNode(val)
		}
	}
}

func (n *Node) FindMedian() float64 {
	size := n.Size
	if size%2 == 0 {
		// even
		a := size / 2
		b := a + 1
		as := n.FindKthSmallest(a)
		bs := n.FindKthSmallest(b)
		return float64(as+bs) / 2
	} else {
		// odd
		mid := size/2 + 1
		return float64(n.FindKthSmallest(mid))
	}
}

// FindKthSmallest
// Positions 1 to L → in left subtree
// Position L+1 → current node
// Positions L+2 and beyond → in right subtree
func (n *Node) FindKthSmallest(kth int) int {
	leftSize := n.getLeftSize()

	if leftSize >= kth {
		return n.Left.FindKthSmallest(kth)
	} else if leftSize+1 == kth {
		return n.Val
	} else { // only could be right side
		return n.Right.FindKthSmallest(kth - leftSize - 1)
	}
}
