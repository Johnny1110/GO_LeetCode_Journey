package word_search

import "testing"

func TestExist(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		word     string
		expected bool
	}{
		{
			name: "Example 1 - ABCCED found",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCCED",
			expected: true,
		},
		{
			name: "Example 2 - SEE found",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "SEE",
			expected: true,
		},
		{
			name: "Example 3 - ABCB not found",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCB",
			expected: false,
		},
		{
			name: "Single cell match",
			board: [][]byte{
				{'A'},
			},
			word:     "A",
			expected: true,
		},
		{
			name: "Single cell no match",
			board: [][]byte{
				{'A'},
			},
			word:     "B",
			expected: false,
		},
		{
			name: "Word longer than board cells",
			board: [][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			word:     "ABCDA",
			expected: false,
		},
		{
			name: "Full board path",
			board: [][]byte{
				{'A', 'B'},
				{'D', 'C'},
			},
			word:     "ABCD",
			expected: true,
		},
		{
			name: "Same letter cannot be reused",
			board: [][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			word:     "ABDA",
			expected: false,
		},
		{
			name: "Snake path",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'E', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCEFSADEESE",
			expected: false,
		},
		{
			name: "Zigzag path",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'E', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCESEEFDS",
			expected: false,
		},
		{
			name: "All same letters - found",
			board: [][]byte{
				{'A', 'A', 'A'},
				{'A', 'A', 'A'},
				{'A', 'A', 'A'},
			},
			word:     "AAAAAAAAA",
			expected: true,
		},
		{
			name: "All same letters - too long",
			board: [][]byte{
				{'A', 'A', 'A'},
				{'A', 'A', 'A'},
				{'A', 'A', 'A'},
			},
			word:     "AAAAAAAAAA",
			expected: false,
		},
		{
			name: "Start from different positions",
			board: [][]byte{
				{'C', 'A', 'A'},
				{'A', 'A', 'A'},
				{'B', 'C', 'D'},
			},
			word:     "AAB",
			expected: true,
		},
		{
			name: "Backtracking required",
			board: [][]byte{
				{'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A'},
			},
			word:     "AAAAAAAAAAAAA",
			expected: false,
		},
		{
			name: "Single row",
			board: [][]byte{
				{'A', 'B', 'C', 'D', 'E'},
			},
			word:     "EDCBA",
			expected: true,
		},
		{
			name: "Single column",
			board: [][]byte{
				{'A'},
				{'B'},
				{'C'},
				{'D'},
			},
			word:     "DCBA",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := exist(tt.board, tt.word)
			if result != tt.expected {
				t.Errorf("exist(%v, %q) = %v; want %v", tt.board, tt.word, result, tt.expected)
			}
		})
	}
}

func BenchmarkExist(b *testing.B) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	benchmarks := []struct {
		name string
		word string
	}{
		{"Found_short", "SEE"},
		{"Found_long", "ABCCED"},
		{"Not_found", "ABCB"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				exist(board, bm.word)
			}
		})
	}
}
