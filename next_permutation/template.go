package next_permutation

func nextPermutation(nums []int) {

	endIdx := len(nums) - 1

	// Step 1. identify the num need to swapped. (pointerA)
	pointerA := -1

	for i := endIdx; i > 0; i-- {
		if nums[i] > nums[i-1] {
			pointerA = i - 1
			break
		}
	}

	// if can not find the pointer, that means input nums is biggest permutation, reverse it and return.
	if pointerA == -1 {
		reverse(nums, 0, endIdx)
		return
	}

	// Step 2. find the element just larger then nums[pointerA] from the right side.
	pointerB := 0

	for i := endIdx; i > pointerA; i-- {
		if nums[i] > nums[pointerA] {
			pointerB = i
			break
		}
	}

	// step 3. swap nums[pointerA] and pointer[B]
	swap(nums, pointerA, pointerB)

	// step 4. reverse the sub array from pointer+1 to the end.
	reverse(nums, pointerA+1, endIdx)
}

func reverse(nums []int, a int, b int) {
	for a < b {
		swap(nums, a, b)
		a++
		b--
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func printNums(nums []int) {
	for i := 0; i < len(nums); i++ {
		print(nums[i])
	}
	println()
}

func Go() {
	input1 := []int{2, 4, 3}
	nextPermutation(input1)
	printNums(input1)

	println()

	input2 := []int{2, 3, 1}
	nextPermutation(input2)
	printNums(input2)

	println()

	input3 := []int{3, 2, 1}
	nextPermutation(input3)
	printNums(input3)

	println()

	input4 := []int{5, 4, 7, 5, 3, 2}
	nextPermutation(input4)
	printNums(input4)
}
