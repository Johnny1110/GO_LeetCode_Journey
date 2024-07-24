package letter_combinations_of_a_phone_number

import "fmt"

var digits = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(input string) []string {
	if input == "" {
		return []string{}
	}
	result := make([]string, 0)
	backtrack(&result, input, "", 0)
	return result
}

func backtrack(result *[]string, input string, temp string, index int) {
	if index >= len(input) {
		*result = append(*result, temp)
		return
	}

	// I will only explain layer-1:

	//in layer-1 user input digit is "1"
	digit := string(input[index])
	// digitString is "abc"
	digitString := digits[digit]

	// for digitString: a, b, c
	for i := 0; i < len(digitString); i++ {
		// in layer-1 temp will be like "a__", "b__", "c__"
		temp = temp + string(digitString[i])
		// go deeper layer-2 to find out "a?_"
		backtrack(result, input, temp, index+1)
		// reset temp (this is called backtrack!)
		temp = temp[:len(temp)-1]
	}
}

func Go() {
	fmt.Println("ans: ", letterCombinations("34"))
}
