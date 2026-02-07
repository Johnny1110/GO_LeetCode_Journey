package find_median_from_data_stream

type MedianFinder struct {
	data []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		data: make([]int, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.data) == 0 {
		this.data = append(this.data, num)
		return
	}

	// find the insert position idx:
	left, right := 0, len(this.data)-1 // init start end.
	positionIdx := this.binarySearch(num, left, right)

	// perform insert:
	this.data = append(this.data, 0) // extend 1 slot.
	copy(this.data[positionIdx+1:], this.data[positionIdx:])
	this.data[positionIdx] = num
}

func (this *MedianFinder) FindMedian() float64 {
	dataLen := len(this.data)
	// odd or even:
	isOdd := dataLen%2 == 1

	if isOdd {
		return float64(this.data[dataLen/2])
	} else {
		// even
		a := this.data[dataLen/2]
		b := this.data[dataLen/2-1]
		return float64(a+b) / 2
	}
}

func (this *MedianFinder) binarySearch(num, left, right int) int {
	if left >= right {
		if this.data[left] > num {
			return left
		} else {
			return left + 1
		}
	}

	middle := (right + left) / 2 // calculate middle

	if this.data[middle] == num {
		return middle + 1
	} else if this.data[middle] > num {
		return this.binarySearch(num, left, middle-1)
	} else {
		return this.binarySearch(num, middle+1, right)
	}
}
