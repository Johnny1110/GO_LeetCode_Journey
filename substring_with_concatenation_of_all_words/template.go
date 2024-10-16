package substring_with_concatenation_of_all_words

import "fmt"

func findSubstring(s string, words []string) []int {
	result := make([]int, 0)

	wordMapProto := make(map[string]int)
	for _, word := range words {
		wordMapProto[word]++
	}

	stepSize := len(words[0])

	for startRowIdx := 0; startRowIdx < stepSize; startRowIdx++ {
		for j := startRowIdx; j+stepSize*len(words) <= len(s); j += stepSize {
			wordMap := make(map[string]int)
			for key, value := range wordMapProto {
				wordMap[key] = value
			}

			if validate(wordMap, s[j:j+stepSize*len(words)], stepSize) {
				result = append(result, j)
			}
		}
	}

	return result
}

func validate(wordMap map[string]int, s string, stepSize int) bool {
	for i := 0; i <= len(s)-stepSize; i = i + stepSize {
		target := s[i : i+stepSize]
		if count, exists := wordMap[target]; !exists || count == 0 {
			return false
		}
		wordMap[target]--
	}

	//println("----------")
	//println("s:", s)
	//println("stepSize:", stepSize)
	//for k, v := range wordMap {
	//	println("MAP K:", k, "V:", v)
	//}
	//
	//println("----------")

	return true
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
