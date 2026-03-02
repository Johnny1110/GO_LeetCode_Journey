package unique_paths

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{
			name: "1x1 grid",
			m:    1,
			n:    1,
			want: 1,
		},
		{
			name: "1x2 grid",
			m:    1,
			n:    2,
			want: 1,
		},
		{
			name: "2x1 grid",
			m:    2,
			n:    1,
			want: 1,
		},
		{
			name: "2x2 grid",
			m:    2,
			n:    2,
			want: 2,
		},
		{
			name: "3x2 grid",
			m:    3,
			n:    2,
			want: 3,
		},
		{
			name: "3x7 grid (LeetCode example 1)",
			m:    3,
			n:    7,
			want: 28,
		},
		{
			name: "3x3 grid (LeetCode example 2)",
			m:    3,
			n:    3,
			want: 6,
		},
		{
			name: "4x4 grid",
			m:    4,
			n:    4,
			want: 20,
		},
		{
			name: "7x3 grid (reverse of example 1)",
			m:    7,
			n:    3,
			want: 28,
		},
		{
			name: "large grid 10x10",
			m:    10,
			n:    10,
			want: 48620,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniquePaths(tt.m, tt.n)
			if got != tt.want {
				t.Errorf("uniquePaths(%d, %d) = %d, want %d", tt.m, tt.n, got, tt.want)
			}
		})
	}
}

func TestUniquePathsEdgeCases(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{
			name: "minimum valid input (1,1)",
			m:    1,
			n:    1,
			want: 1,
		},
		{
			name: "single row",
			m:    1,
			n:    10,
			want: 1,
		},
		{
			name: "single column",
			m:    10,
			n:    1,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniquePaths(tt.m, tt.n)
			if got != tt.want {
				t.Errorf("uniquePaths(%d, %d) = %d, want %d", tt.m, tt.n, got, tt.want)
			}
		})
	}
}
