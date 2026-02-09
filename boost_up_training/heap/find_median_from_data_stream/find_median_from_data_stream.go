package find_median_from_data_stream

type MedianFinder struct {
	treeRoot *Node
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	if this.treeRoot == nil {
		this.treeRoot = NewNode(num)
		return
	}

	this.treeRoot.Push(num)
}

func (this *MedianFinder) FindMedian() float64 {
	return this.treeRoot.FindMedian()
}
