package bucket_sort_solutation

func topKFrequent(nums []int, k int) []int {
	// make a frequency map
	numFreqMap := make(map[int]int)
	for _, num := range nums {
		numFreqMap[num]++
	}

	// make a bucket sort map, max is len(nums)

	buckets := make([][]int, len(nums)+1)

	// put freqMap into bucket:
	for num, freq := range numFreqMap {
		buckets[freq] = append(buckets[freq], num)
	}

	result := make([]int, 0, k)
	for i := len(buckets) - 1; i > 0 && len(result) < k; i-- {
		result = append(result, buckets[i]...)
	}

	return result[:k]
}
