package find_median_from_data_stream

type AVLNode struct {
	Val    int
	Left   *AVLNode
	Right  *AVLNode
	Height int // track height instead of size
	Size   int // keep for FindKth
}

func NewAVLNode(x int) *AVLNode {
	return &AVLNode{
		Val:    x,
		Size:   1,
		Height: 1,
	}
}

// getBalanceFactor
// BF > 1 → Left heavy
// BF < -1 → Right heavy
func (n *AVLNode) getBalanceFactor() int {
	return n.GetLeftHeight() - n.GetRightHeight()
}

// Push return height of current node
// return new root
func (n *AVLNode) Push(x int) *AVLNode {
	// 1. insert as normal BST, update Height
	n.Size++
	if n.Val >= x {
		// goes left
		if n.Left == nil {
			n.Left = NewAVLNode(x)
		} else {
			n.Left = n.Left.Push(x)
		}
	} else {
		// goes right
		if n.Right == nil {
			n.Right = NewAVLNode(x)
		} else {
			n.Right = n.Right.Push(x)
		}
	}

	// 2. Update THIS node's height
	n.Height = n.calculateHeight()

	// 3. check the balance factor
	balanceFactor := n.getBalanceFactor()
	// factor > 1: left heavy
	// factor == 0: balance
	// factor < -1: right heavy

	// select four cases: <LL, RR, LR, RL>
	if balanceFactor > 1 { // LL or LR
		if n.Left.getBalanceFactor() >= 0 {
			// LL (single right rotation)
			return rightRotate(n)
		} else {
			// LR (LEFT rotate left child, then RIGHT rotate root)
			n.Left = leftRotate(n.Left)
			return rightRotate(n)
		}
	} else if balanceFactor < -1 { // // RR or RL
		if n.Right.getBalanceFactor() <= 0 {
			// RR (single left rotation)
			return leftRotate(n)
		} else {
			// RL (RIGHT rotate right child, then LEFT rotate root)
			n.Right = rightRotate(n.Right)
			return leftRotate(n)
		}
	} else {
		// balance, just return
		return n
	}
}

func (n *AVLNode) GetLeftHeight() int {
	if n.Left == nil {
		return 0
	}
	return n.Left.Height
}

func (n *AVLNode) GetRightHeight() int {
	if n.Right == nil {
		return 0
	}
	return n.Right.Height
}

func (n *AVLNode) calculateHeight() int {
	return max(n.GetRightHeight(), n.GetLeftHeight()) + 1
}

func (n *AVLNode) GetLeftSize() int {
	if n.Left == nil {
		return 0
	}
	return n.Left.Size
}

func (n *AVLNode) GetRightSize() int {
	if n.Right == nil {
		return 0
	}
	return n.Right.Size
}

// rightRotate
//
//	   y                    x
//	  / \                  / \
//	 x   C    →           A   y
//	/ \                      / \
//
// A   B                    B   C
// ========================================
func rightRotate(y *AVLNode) *AVLNode {
	x := y.Left
	B := x.Right

	// Perform rotation
	x.Right = y
	y.Left = B

	// Update heights (y first, because x depends on y now)
	// not changed: C, B A
	// changed: y, x
	y.Height = y.calculateHeight()
	x.Height = x.calculateHeight()

	// Update sizes
	y.Size = y.GetLeftSize() + y.GetRightSize() + 1
	x.Size = x.GetLeftSize() + x.GetRightSize() + 1

	return x
}

// leftRotate
//
//	  x                      y
//	 / \                    / \
//	A   y      →           x   C
//	   / \                / \
//	  B   C              A   B
func leftRotate(x *AVLNode) *AVLNode {
	y := x.Right
	B := y.Left

	y.Left = x
	x.Right = B

	// update height
	x.Height = x.calculateHeight()
	y.Height = y.calculateHeight()

	// update size
	x.Size = x.GetLeftSize() + x.GetRightSize() + 1
	y.Size = y.GetLeftSize() + y.GetRightSize() + 1

	return y
}

// FindKthSmallest
// Positions 1 to L → in left subtree
// Position L+1 → current node
// Positions L+2 and beyond → in right subtree
func (n *AVLNode) FindKthSmallest(kth int) int {
	leftSize := n.GetLeftSize()

	if leftSize >= kth {
		return n.Left.FindKthSmallest(kth)
	} else if leftSize+1 == kth {
		return n.Val
	} else { // only could be right side
		return n.Right.FindKthSmallest(kth - leftSize - 1)
	}
}
