package substring_with_concatenation_of_all_words_self_try

import "fmt"

// def string queue --------------------------------------------

type StringQueue struct {
	items []string
}

func (q *StringQueue) Enqueue(item string) {
	q.items = append(q.items, item)
}

func (q *StringQueue) Dequeue() string {
	if len(q.items) == 0 {
		panic("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *StringQueue) Size() int {
	return len(q.items)
}

func (q *StringQueue) Reset() {
	q.items = q.items[:0]
}

func (q *StringQueue) Print() {
	for _, v := range q.items {
		fmt.Print(v)
	}
	fmt.Print("\n")
}

// --------------------------------------------

func findSubstring(s string, words []string) []int {

	wordLen := len(words[0])
	wordListLen := len(words)
	fullWordLen := wordLen * wordListLen
	wordMap := make(map[string]int)

	result := []int{}

	// init wordMap
	for _, val := range words {
		wordMap[val]++
	}
	// init monitor queue
	queue := &StringQueue{}
	// init seenTime map
	seenTimes := make(map[string]int)

	reset := func() {
		queue.Reset()
		for k := range seenTimes {
			delete(seenTimes, k)
		}
	}

	for i := 0; i < wordLen; i++ {

		reset()

		for j := i; j+wordLen <= len(s); j += wordLen {

			currentWord := s[j : j+wordLen]

			currentWordCount, exists := wordMap[currentWord]

			if !exists {
				reset()
				continue
			}

			for currentWordCount <= seenTimes[currentWord] {
				discardWord := queue.Dequeue()
				seenTimes[discardWord]--
			}

			queue.Enqueue(currentWord)
			seenTimes[currentWord]++

			if queue.Size() == wordListLen {
				//queue.Print()
				result = append(result, j+wordLen-fullWordLen)
			}
		}

	}

	return result
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
