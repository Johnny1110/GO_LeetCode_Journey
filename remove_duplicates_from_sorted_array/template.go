package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	pointerA := 0

	for pointerB := 1; pointerB < len(nums); pointerB++ {

		if nums[pointerA] != nums[pointerB] {
			nums[pointerA+1] = nums[pointerB]
			pointerA++
		}
	}

	return pointerA + 1
}

// Go call this func in main.go
func Go() {
	input1 := []int{1, 2, 3, 3, 5, 6, 6}
	ans := removeDuplicates(input1)
	println("answer: 5, result:", ans)

	input2 := []int{1, 2, 2, 3, 3, 6, 6}
	ans = removeDuplicates(input2)
	println("answer: 4, result:", ans)

	input3 := []int{1, 1, 1, 3, 3, 3, 6}
	ans = removeDuplicates(input3)
	println("answer: 3, result:", ans)
}
