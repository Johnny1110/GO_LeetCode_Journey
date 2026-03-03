package course_schedule

// example: prerequisites := [[0, 1], [1, 0]]
func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	// 1. create  directed graph
	dirgraph := make(map[int]IntSet)

	for _, row := range prerequisites {
		courseA, courseB := row[0], row[1]

		// build dirgraph
		if _, exists := dirgraph[courseB]; !exists {
			dirgraph[courseB] = NewIntSet()
		}

		dirgraph[courseB].Add(courseA)
	}

	stateTracking := make([]int, numCourses)	// idx: course, val: state

	// state:
	// 0: non-visited
	// 1: in_progress
	// 2: compeleted
	
	var dfs func(course int) bool
	dfs = func(course int) bool {
		if stateTracking[course] == 1 {
			return false
		}else if stateTracking[course] == 2 {
			return true
		} else {
			// 0: non-visited
			stateTracking[course] = 1 // mark this course as in_progress

			if courseSet, exists := dirgraph[course]; exists {
				for otherCourse, _ := range courseSet {
					if !dfs(otherCourse) {
						return false
					}
				}
			}

			// expolore complelted
			stateTracking[course] = 2
			return true
		}
	}

	for course := range numCourses {
		if !dfs(course) {
			return false
		}
	}

	return true
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
