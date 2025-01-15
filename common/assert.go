package common

import (
	"fmt"
	"reflect"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

var WRONG = "[WRONG!]"
var GOOD = "[PASSED!]"

func Assert_answer_or(input any, expects ...any) {
	flag := false
	for _, expect := range expects {
		if assert(input, expect) {
			flag = true
			break
		}
	}
	if flag {
		fmt.Println("* " + Green + GOOD + " ans: " + fmt.Sprint(input) + " expect in: " + fmt.Sprint(expects) + Reset)
	} else {
		fmt.Println("* " + Red + WRONG + " ans: " + fmt.Sprint(input) + " expect: " + fmt.Sprint(expects) + Reset)
	}
}

func Assert_answer(input any, expect any) {
	if !assert(input, expect) {
		fmt.Println("* " + Red + WRONG + " ans: " + fmt.Sprint(input) + " expect: " + fmt.Sprint(expect) + Reset)
	} else {
		fmt.Println("* " + Green + GOOD + " ans: " + fmt.Sprint(input) + " expect: " + fmt.Sprint(expect) + Reset)
	}
}

func assert(input any, expect any) bool {
	// Check if both input and expect are slices of integers
	if reflect.TypeOf(input) == reflect.TypeOf([]int{}) && reflect.TypeOf(expect) == reflect.TypeOf([]int{}) {
		// Use reflect.DeepEqual to compare slices
		return reflect.DeepEqual(input, expect)
	}

	return input == expect
}
