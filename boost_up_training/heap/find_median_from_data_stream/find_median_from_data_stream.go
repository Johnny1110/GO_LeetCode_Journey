package find_median_from_data_stream

type MedianFinder struct {
	treeRoot *AVLNode
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	if this.treeRoot == nil {
		this.treeRoot = NewAVLNode(num)
		return
	}

	this.treeRoot = this.treeRoot.Push(num)
}

func (this *MedianFinder) FindMedian() float64 {
	n := this.treeRoot
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
