package course_schedule

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	//Input: numCourses = 2, prerequisites = [[1,0]]
	//Output: true
	//Explanation: There are a total of 2 courses to take.
	//To take course 1 you should have finished course 0. So it is possible.
	prerequisites := [][]int{
		{1, 0},
	}
	assert.True(t, canFinish(2, prerequisites))
}

func Test_2(t *testing.T) {
	//Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
	//Output: false
	//Explanation: There are a total of 2 courses to take.
	//To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.
	//
	prerequisites := [][]int{
		{1, 0},
		{0, 1},
	}
	assert.False(t, canFinish(2, prerequisites))
}
