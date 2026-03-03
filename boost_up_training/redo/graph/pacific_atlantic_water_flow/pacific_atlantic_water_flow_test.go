package pacific_atlantic_water_flow

import (
	"sort"
	"testing"
)

// normalizeResult sorts the result so we can compare regardless of order.
// Why: the problem says "return in any order", so our solution might
// produce [0,4] before [0,3] or vice versa. Sorting both expected
// and actual lets us use a simple deep-equal comparison.
func normalizeResult(res [][]int) [][]int {
	sort.Slice(res, func(i, j int) bool {
		if res[i][0] != res[j][0] {
			return res[i][0] < res[j][0]
		}
		return res[i][1] < res[j][1]
	})
	return res
}

func equal2D(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestPacificAtlantic(t *testing.T) {
	tests := []struct {
		name     string
		heights  [][]int
		expected [][]int
	}{
		// ---------------------------------------------------------------
		// Case 1: Classic example from LeetCode
		// Why: baseline correctness check; the expected output is known.
		// The grid has peaks at corners and a ridge through the middle.
		// ---------------------------------------------------------------
		{
			name: "classic_example",
			heights: [][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			expected: [][]int{
				{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0},
			},
		},

		// ---------------------------------------------------------------
		// Case 2: Single cell (1x1)
		// Why: the smallest possible grid. The single cell borders both
		// oceans simultaneously, so it must appear in the result.
		// Tests boundary initialization when m=1, n=1.
		// ---------------------------------------------------------------
		{
			name:     "single_cell",
			heights:  [][]int{{1}},
			expected: [][]int{{0, 0}},
		},

		// ---------------------------------------------------------------
		// Case 3: Single row (1xN)
		// Why: every cell in a single row touches Pacific (top edge) and
		// Atlantic (bottom edge — since row 0 is also the last row).
		// Left edge is Pacific, right edge is Atlantic, and the single
		// row IS the top edge (Pacific) AND the bottom edge (Atlantic).
		// So every cell reaches both.
		// ---------------------------------------------------------------
		{
			name:    "single_row",
			heights: [][]int{{1, 2, 3, 4, 5}},
			expected: [][]int{
				{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4},
			},
		},

		// ---------------------------------------------------------------
		// Case 4: Single column (Nx1)
		// Why: mirror of single row. Tests that the code handles m>1, n=1
		// correctly. Every cell is on both left (Pacific) and right
		// (Atlantic) edges.
		// ---------------------------------------------------------------
		{
			name:    "single_column",
			heights: [][]int{{1}, {2}, {3}, {4}, {5}},
			expected: [][]int{
				{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0},
			},
		},

		// ---------------------------------------------------------------
		// Case 5: All same height
		// Why: when every cell has equal elevation, water can flow in
		// every direction (>=). So every cell reaches both oceans.
		// Catches bugs where the code uses strict '>' instead of '>='.
		// ---------------------------------------------------------------
		{
			name: "uniform_height",
			heights: [][]int{
				{3, 3, 3},
				{3, 3, 3},
				{3, 3, 3},
			},
			expected: [][]int{
				{0, 0}, {0, 1}, {0, 2},
				{1, 0}, {1, 1}, {1, 2},
				{2, 0}, {2, 1}, {2, 2},
			},
		},

		// ---------------------------------------------------------------
		// Case 6: Strictly decreasing top-left to bottom-right
		// Why: water flows naturally toward bottom-right (Atlantic), but
		// only edge cells can reach Pacific via the edge itself.
		// Inner cells can reach Atlantic but NOT Pacific because every
		// neighbor toward top/left is strictly higher.
		//
		// Pacific-reachable (BFS from top/left edges going to >= height):
		//   top row + left col only, since inner cells are lower.
		// Atlantic-reachable (BFS from bottom/right edges going to >=):
		//   bottom row + right col only.
		// Intersection: cells on BOTH edges = the 3 corners that touch both.
		// ---------------------------------------------------------------
		{
			name: "decreasing_diagonal",
			heights: [][]int{
				{5, 4, 3},
				{4, 3, 2},
				{3, 2, 1},
			},
			expected: [][]int{
				{0, 0}, {0, 2}, {2, 0},
			},
		},

		// ---------------------------------------------------------------
		// Case 7: Strictly increasing top-left to bottom-right
		// Why: mirror of case 6. Now water flows easily toward Pacific
		// (top/left), but reaching Atlantic requires being on the
		// bottom/right edges.
		// ---------------------------------------------------------------
		{
			name: "increasing_diagonal",
			heights: [][]int{
				{1, 2, 3},
				{2, 3, 4},
				{3, 4, 5},
			},
			expected: [][]int{
				{0, 2}, {2, 0}, {2, 2},
			},
		},

		// ---------------------------------------------------------------
		// Case 8: High border, low center (valley)
		// Why: the border cells all touch an ocean edge and are high
		// enough to drain to each other along the border. The center
		// is a valley — water from center (1) can't climb to any
		// neighbor (all 5), so center does NOT reach either ocean.
		// Tests that reverse BFS from ocean edges doesn't incorrectly
		// descend into a valley.
		// ---------------------------------------------------------------
		{
			name: "valley_center",
			heights: [][]int{
				{5, 5, 5},
				{5, 1, 5},
				{5, 5, 5},
			},
			expected: [][]int{
				{0, 0}, {0, 1}, {0, 2},
				{1, 0}, {1, 2},
				{2, 0}, {2, 1}, {2, 2},
			},
		},

		// ---------------------------------------------------------------
		// Case 9: Single peak in center
		// Why: center (9) is the highest point, clearly drains to both.
		// All border cells are 1 — equal height, so water flows freely
		// along the entire border, meaning every border cell reaches
		// both oceans too. Result: all 9 cells.
		// This verifies that equal-height traversal along edges works.
		// ---------------------------------------------------------------
		{
			name: "peak_center",
			heights: [][]int{
				{1, 1, 1},
				{1, 9, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{0, 0}, {0, 1}, {0, 2},
				{1, 0}, {1, 1}, {1, 2},
				{2, 0}, {2, 1}, {2, 2},
			},
		},

		// ---------------------------------------------------------------
		// Case 10: 2x2 grid with distinct values
		// Why: smallest non-trivial grid. Manual trace:
		//   Pacific edge: row 0 + col 0 → cells (0,0),(0,1),(1,0)
		//   Atlantic edge: row 1 + col 1 → cells (0,1),(1,0),(1,1)
		//
		//   (0,0)=1: on Pacific. Atlantic? can't flow to (0,1)=2 or (1,0)=3. ✗
		//   (0,1)=2: Pacific ✓ (top). Atlantic ✓ (right). ✓
		//   (1,0)=3: Pacific ✓ (left). Atlantic ✓ (bottom). ✓
		//   (1,1)=4: Atlantic ✓. Pacific? 4→2 (yes, 4≥2), 2 is on Pacific. ✓
		// ---------------------------------------------------------------
		{
			name: "two_by_two",
			heights: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: [][]int{
				{0, 1}, {1, 0}, {1, 1},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := pacificAtlantic(tc.heights)
			got = normalizeResult(got)
			expected := normalizeResult(tc.expected)

			if !equal2D(got, expected) {
				t.Errorf("\nheights:  %v\nexpected: %v\ngot:      %v", tc.heights, expected, got)
			}
		})
	}
}