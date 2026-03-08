package top_k_frequent_elements

func topKFrequent(nums []int, k int) []int {

	numFreq := make(map[int]int)
	for _, num := range nums {
		numFreq[num]++
	}

	freqBucket := make([][]int, len(nums)+1)
	for num, freq := range numFreq {
		freqBucket[freq] = append(freqBucket[freq], num)
	}

	result := make([]int, 0, k*2)

	for freq := len(nums); freq >= 0; freq-- {
		if k == 0 {
			break
		}

		numbers := freqBucket[freq]
		result = append(result, numbers...)
		k -= len(numbers)

	}

	return result
}
