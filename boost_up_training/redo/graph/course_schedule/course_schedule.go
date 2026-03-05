package course_schedule

// example: prerequisites := [[0, 1], [1, 0]]
func canFinish(numCourses int, prerequisites [][]int) bool {

	// init dir graph
	dirgraph := make(map[int]IntSet)
	for course := range numCourses {
		dirgraph[course] = IntSet(make(map[int]struct{}))
	}

	for _, prerequisite := range prerequisites {
		courseA, courceB := prerequisite[0], prerequisite[1]
		dirgraph[courseA].Add(courceB)
	}

	stateTracking := make(map[int]State)

	var dfs func(course int) bool
	dfs = func(course int) bool {
		switch stateTracking[course] {
		case PENDING:
			return false
		case COMPLETED:
			return true
		case UNVISITED:

			stateTracking[course] = PENDING

			for nextCourse, _ := range dirgraph[course] {
				if !dfs(nextCourse) {
					return false
				}
			}
			stateTracking[course] = COMPLETED
			return true
		default:
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

// ==================================
type IntSet map[int]struct{}

func (s IntSet) Add(val int) {
	s[val] = struct{}{}
}

type State int

const (
	UNVISITED State = iota
	PENDING
	COMPLETED
)
