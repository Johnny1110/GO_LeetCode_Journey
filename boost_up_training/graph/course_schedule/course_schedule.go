package course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	// Step 1: Build the adjacency list.
	// key: course, value: required course
	adjacencyList := make(map[int][]int)

	for _, prereq := range prerequisites {
		target := prereq[0]
		depend := prereq[1]
		adjacencyList[target] = append(adjacencyList[target], depend)
	}

	// Step 2: Set up the state tracking.
	// index: course, value: state (0: Unvisited, 1: In progress, 2: Completed)
	stateTracking := make([]int, numCourses)

	// Step 3: Implement the DFS function.
	var dfs func(course int) bool
	dfs = func(course int) bool {
		// 3-1: check this course's visit state.
		switch stateTracking[course] {
		case 0: // "Unvisited"
			stateTracking[course] = 1 // marked as "In progress"
			// 3-2: explore next courses
			for _, c := range adjacencyList[course] {
				if !dfs(c) {
					// should be all success
					return false
				}
			}

			// all course is good, mark this course as completed
			stateTracking[course] = 2
			return true

		case 1: // "In progress"
			return false
		case 2: // "Completed"
			return true
		default:
			return false
		}
	}

	// Step 4: Put it all together.
	for c := range numCourses {
		if !dfs(c) {
			return false
		}
	}

	return true
}
