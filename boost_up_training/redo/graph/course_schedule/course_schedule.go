package course_schedule

import "fmt"

// example: prerequisites := [[0, 1], [1, 0]]
func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	dependency := make(map[int]int)

	// 1. create  directed graph
	dirgraph := make(map[int]IntSet)

	for _, row := range prerequisites {
		courseA, courseB := row[0], row[1]

		// init dependency if needed
		if _, exists := dependency[courseA]; exists {
			dependency[courseA]++
		} else {
			dependency[courseA] = 1
		}

		if _, exists := dependency[courseB]; !exists {
			dependency[courseB] = 0
		}

		// build dirgraph
		if _, exists := dirgraph[courseB]; !exists {
			dirgraph[courseB] = NewIntSet()
		}

		dirgraph[courseB].Add(courseA)
	}

	// must find a course not in any set.
	for head, v := range dependency {
		if v == 0 { // not depended by other course
			fmt.Printf("head: %v \n", head)
		}
	}

	return false
}

// ======================================================

type IntSet map[int]struct{}

func (s IntSet) Len() int {
	return len(s)
}

func (s IntSet) Add(val int) {
	s[val] = struct{}{}
}

func NewIntSet() IntSet {
	return IntSet(make(map[int]struct{}))
}
