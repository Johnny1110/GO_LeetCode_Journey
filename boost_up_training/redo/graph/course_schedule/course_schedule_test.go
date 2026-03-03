package course_schedule

import "testing"

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		expected      bool
	}{
		{
			name:          "no courses",
			numCourses:    0,
			prerequisites: [][]int{},
			expected:      true,
		},
		{
			name:          "single course no prerequisites",
			numCourses:    1,
			prerequisites: [][]int{},
			expected:      true,
		},
		{
			name:          "two courses no prerequisites",
			numCourses:    2,
			prerequisites: [][]int{},
			expected:      true,
		},
		{
			name:          "LeetCode example 1 - valid schedule",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			expected:      true,
		},
		{
			name:          "Example 1",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			expected:      false,
		},
		{
			name:          "Example 2",
			numCourses:    3,
			prerequisites: [][]int{{1, 0}, {2, 1}},
			expected:      true,
		},
		{
			name:          "Example 3",
			numCourses:    3,
			prerequisites: [][]int{{1, 0}, {2, 1}, {0, 2}},
			expected:      false,
		},

		{
			name:          "Example 4",
			numCourses:    5,
			prerequisites: [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}},
			expected:      true,
		},

		{
			name:          "complex valid schedule",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			expected:      true,
		},
		{
			name:          "multiple independent chains",
			numCourses:    6,
			prerequisites: [][]int{{1, 0}, {2, 1}, {4, 3}, {5, 4}},
			expected:      true,
		},
		{
			name:          "larger cycle",
			numCourses:    5,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}, {4, 3}, {0, 4}},
			expected:      false,
		},
		{
			name:          "self dependency",
			numCourses:    2,
			prerequisites: [][]int{{0, 0}},
			expected:      false,
		},
		{
			name:          "multiple prerequisites for one course",
			numCourses:    4,
			prerequisites: [][]int{{2, 0}, {2, 1}, {3, 2}},
			expected:      true,
		},
		{
			name:          "complex graph with cycle",
			numCourses:    6,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}, {1, 3}, {4, 1}, {5, 4}},
			expected:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canFinish(tt.numCourses, tt.prerequisites)
			if result != tt.expected {
				t.Errorf("canFinish(%d, %v) = %v, want %v",
					tt.numCourses, tt.prerequisites, result, tt.expected)
			}
		})
	}
}

func TestCanFinishEdgeCases(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		expected      bool
	}{
		{
			name:          "zero courses",
			numCourses:    0,
			prerequisites: [][]int{},
			expected:      true,
		},
		{
			name:          "large number of courses no dependencies",
			numCourses:    100,
			prerequisites: [][]int{},
			expected:      true,
		},
		{
			name:          "single course with self-reference",
			numCourses:    1,
			prerequisites: [][]int{{0, 0}},
			expected:      false,
		},
		{
			name:          "empty prerequisites array",
			numCourses:    5,
			prerequisites: [][]int{},
			expected:      true,
		},
		{
			name:          "prerequisite references non-existent course",
			numCourses:    2,
			prerequisites: [][]int{{1, 3}},
			expected:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canFinish(tt.numCourses, tt.prerequisites)
			if result != tt.expected {
				t.Errorf("canFinish(%d, %v) = %v, want %v",
					tt.numCourses, tt.prerequisites, result, tt.expected)
			}
		})
	}
}

func TestCanFinishCycleDetection(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		expected      bool
		description   string
	}{
		{
			name:          "simple 2-node cycle",
			numCourses:    2,
			prerequisites: [][]int{{0, 1}, {1, 0}},
			expected:      false,
			description:   "Course 0 depends on 1, and 1 depends on 0",
		},
		{
			name:          "3-node cycle",
			numCourses:    3,
			prerequisites: [][]int{{0, 1}, {1, 2}, {2, 0}},
			expected:      false,
			description:   "0->1->2->0 forms a cycle",
		},
		{
			name:          "4-node cycle with extra edge",
			numCourses:    4,
			prerequisites: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {2, 1}},
			expected:      false,
			description:   "Contains cycle 0->1->2->3->0",
		},
		{
			name:          "no cycle - tree structure",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}},
			expected:      true,
			description:   "Tree structure with root at course 0",
		},
		{
			name:          "no cycle - DAG",
			numCourses:    5,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 1}, {4, 2}, {4, 3}},
			expected:      true,
			description:   "Valid DAG structure",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canFinish(tt.numCourses, tt.prerequisites)
			if result != tt.expected {
				t.Errorf("canFinish(%d, %v) = %v, want %v. %s",
					tt.numCourses, tt.prerequisites, result, tt.expected, tt.description)
			}
		})
	}
}
