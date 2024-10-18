package substring_with_concatenation_of_all_words_ans

import "fmt"

func findSubstring(s string, words []string) []int {
	wordMapCount := make(map[string]int)
	lenW := len(words[0])
	lenWs := len(words)
	lenP := lenW * lenWs
	lenS := len(s)
	arrRes := make([]int, 0)

	// init count of each word
	for i := 0; i < len(words); i++ {
		wordMapCount[words[i]]++
	}

	queue := make([]string, 0, lenWs) // queue to evaluate sliding window
	seenTimes := make(map[string]int) // how many times a word has been seen

	head, queueLen := 0, 0 // we use pointer instead of choping the queue head every time

	reset := func() {
		head, queueLen = 0, 0
		queue = queue[:0]          // apparently this is the fastest reinit a slice with the same size
		for k := range seenTimes { // to reinit the map we delete each key of the map
			delete(seenTimes, k)
		}
	}

	for i := 0; i < lenW; i++ { //sliding window starting from each char in fisrt word length
		reset()
		for j := i; j+lenW <= lenS; j += lenW {
			curWord := s[j : j+lenW]

			curWordCount, exist := wordMapCount[curWord]
			if !exist { // if curword not exist in words reset the queue
				reset()
				continue
			}

			for seenTimes[curWord] >= curWordCount { // if the word has been seen as much as the count, chop the head of queue until the seen times is less then
				seenTimes[queue[head]]-- // decrease the seen times of the head
				head++                   // next index of the queue became the head
				queueLen--               // decrease the queue len
			}

			//append that word
			seenTimes[curWord]++
			queue = append(queue, curWord)
			queueLen++

			if queueLen == lenWs { // all words was found subsquencely
				arrRes = append(arrRes, j+lenW-lenP) // the index of the words combination is the first index of the queue head
			}
		}
	}

	return arrRes
}

// Go call this func in main.go
func Go() {
	ans1 := findSubstring("barfoothefoobarman", []string{"foo", "bar"})
	fmt.Println("ans1:", ans1)

	ans2 := findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"})
	fmt.Println("ans2:", ans2)

	ans3 := findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"})
	fmt.Println("ans3:", ans3)

	// "lingmindraboofooowingdingbarrwingmonkeypoundcake"
	// ["fooo","barr","wing","ding","wing"]
	ans4 := findSubstring("lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"})
	fmt.Println("ans4:", ans4)

}
