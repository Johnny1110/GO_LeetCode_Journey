package find_median_from_data_stream

type MedianFinder struct {
	nums []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		nums: []int{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	// binsearch position
	targetIdx := binarySearch(this.nums, num)

	if targetIdx == len(this.nums) {
		this.nums = append(this.nums, num)
	} else {
		// insert
		this.nums = append(this.nums[:targetIdx+1], this.nums[targetIdx:]...)
		this.nums[targetIdx] = num
	}

}

func (this *MedianFinder) FindMedian() float64 {
	length := len(this.nums)
	if length%2 == 0 {
		// even
		numA := this.nums[length/2]
		numB := this.nums[(length/2)-1]
		return float64(numA+numB) / 2
	} else {
		// odd
		return float64(this.nums[length/2])
	}
}

func binarySearch(nums []int, num int) int {

	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == num {
			return mid + 1
		} else if nums[mid] > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
