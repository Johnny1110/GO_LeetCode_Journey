package coin_change

import (
	"math"
	"testing"
)

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{
		{
			name:     "Example 1",
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3, // 5+5+1
		},
		{
			name:     "Example 2 - [2] amount=3",
			coins:    []int{2},
			amount:   3,
			expected: -1,
		},
		{
			name:     "Example 3 - [1] amount=0",
			coins:    []int{1},
			amount:   0,
			expected: 0,
		},
		{
			name:     "zero amount with multiple coins",
			coins:    []int{1, 2, 5},
			amount:   0,
			expected: 0,
		},
		{
			name:     "exact match single coin",
			coins:    []int{5},
			amount:   5,
			expected: 1,
		},
		{
			name:     "amount smaller than all coins",
			coins:    []int{5},
			amount:   3,
			expected: -1,
		},
		{
			name:     "odd amount with only even coins",
			coins:    []int{2},
			amount:   1,
			expected: -1,
		},
		{
			name:     "single coin denomination repeated many times",
			coins:    []int{1},
			amount:   5,
			expected: 5,
		},
		{
			name:     "large amount use largest coin only",
			coins:    []int{1, 2, 5},
			amount:   100,
			expected: 20, // 5*20
		},
		{
			name:     "three denominations mixed",
			coins:    []int{2, 5, 10},
			amount:   27,
			expected: 4, // 10+10+5+2
		},
		{
			name:     "two of the same coin needed",
			coins:    []int{3, 7},
			amount:   14,
			expected: 2, // 7+7
		},
		{
			name:     "greedy fails - DP picks 3+3 over 4+1+1",
			coins:    []int{1, 3, 4},
			amount:   6,
			expected: 2, // 3+3=6, greedy would pick 4+1+1=3 coins
		},
		{
			name:     "greedy fails - skip large coin for 6+5",
			coins:    []int{9, 6, 5, 1},
			amount:   11,
			expected: 2, // 6+5=11, greedy would pick 9+1+1=3 coins
		},
		{
			name:     "even coins cannot make odd amount",
			coins:    []int{2, 4},
			amount:   7,
			expected: -1,
		},
		{
			name:     "two different coins sum to amount",
			coins:    []int{1, 5, 6, 9},
			amount:   11,
			expected: 2, // 6+5=11
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := coinChange(tt.coins, tt.amount)
			if result != tt.expected {
				t.Errorf("coinChange(%v, %d) = %d, expected %d", tt.coins, tt.amount, result, tt.expected)
			}
		})
	}
}

func TestCoinChangeEdgeCases(t *testing.T) {
	t.Run("amount is zero", func(t *testing.T) {
		result := coinChange([]int{1, 2, 5}, 0)
		if result != 0 {
			t.Errorf("coinChange([1,2,5], 0) = %d, expected 0", result)
		}
	})

	t.Run("single coin denomination 1 - large amount", func(t *testing.T) {
		result := coinChange([]int{1}, 10000)
		if result != 10000 {
			t.Errorf("coinChange([1], 10000) = %d, expected 10000", result)
		}
	})

	t.Run("coin much larger than max amount", func(t *testing.T) {
		result := coinChange([]int{math.MaxInt32}, 10000)
		if result != -1 {
			t.Errorf("coinChange([MaxInt32], 10000) = %d, expected -1", result)
		}
	})

	t.Run("coin value equals amount exactly", func(t *testing.T) {
		result := coinChange([]int{3, 7, 11}, 11)
		if result != 1 {
			t.Errorf("coinChange([3,7,11], 11) = %d, expected 1", result)
		}
	})

	t.Run("all coins larger than amount", func(t *testing.T) {
		result := coinChange([]int{5, 10, 25}, 3)
		if result != -1 {
			t.Errorf("coinChange([5,10,25], 3) = %d, expected -1", result)
		}
	})

	t.Run("single coin denomination 2 - even amount", func(t *testing.T) {
		result := coinChange([]int{2}, 10000)
		if result != 5000 {
			t.Errorf("coinChange([2], 10000) = %d, expected 5000", result)
		}
	})

	t.Run("coin equals 1 - always solvable for any positive amount", func(t *testing.T) {
		for amount := 1; amount <= 10; amount++ {
			result := coinChange([]int{1, 3, 5}, amount)
			if result == -1 {
				t.Errorf("coinChange([1,3,5], %d) = -1, should be solvable when coin 1 exists", amount)
			}
		}
	})
}

func TestCoinChangeSpecialCases(t *testing.T) {
	t.Run("ascending order coins", func(t *testing.T) {
		result := coinChange([]int{1, 5, 10, 25}, 41)
		if result != 4 { // 25+10+5+1
			t.Errorf("coinChange([1,5,10,25], 41) = %d, expected 4", result)
		}
	})

	t.Run("descending order coins - same result", func(t *testing.T) {
		result := coinChange([]int{25, 10, 5, 1}, 41)
		if result != 4 { // coin order should not affect result
			t.Errorf("coinChange([25,10,5,1], 41) = %d, expected 4", result)
		}
	})

	t.Run("prime amount with prime coins", func(t *testing.T) {
		result := coinChange([]int{3, 5, 7}, 11)
		if result != 3 { // 3+3+5=11
			t.Errorf("coinChange([3,5,7], 11) = %d, expected 3", result)
		}
	})

	t.Run("multiple paths to same min coins", func(t *testing.T) {
		// 10 = 5+5 (2 coins) or 2+2+2+2+2 (5 coins), min is 2
		result := coinChange([]int{2, 5}, 10)
		if result != 2 {
			t.Errorf("coinChange([2,5], 10) = %d, expected 2", result)
		}
	})

	t.Run("impossible - no coin divides amount", func(t *testing.T) {
		result := coinChange([]int{3, 5}, 1)
		if result != -1 {
			t.Errorf("coinChange([3,5], 1) = %d, expected -1", result)
		}
	})

	t.Run("amount is 1 with coin 1", func(t *testing.T) {
		result := coinChange([]int{1}, 1)
		if result != 1 {
			t.Errorf("coinChange([1], 1) = %d, expected 1", result)
		}
	})
}

func TestCoinChangeComplexScenarios(t *testing.T) {
	t.Run("classic DP - skip larger coin for two medium coins", func(t *testing.T) {
		// greedy picks 10, then stuck; DP finds 7+7=14
		result := coinChange([]int{6, 7, 10}, 14)
		if result != 2 {
			t.Errorf("coinChange([6,7,10], 14) = %d, expected 2", result)
		}
	})

	t.Run("many denominations - large amount", func(t *testing.T) {
		coins := []int{1, 2, 5, 10, 20, 50, 100, 200, 500}
		result := coinChange(coins, 9999)
		// 500*19=9500, 200*2=400, 50*1=50, 20*2=40, 5*1=5, 2*2=4, 1*1=1 → 9500+400+50+40+5+4+1=10000 too much
		// 500*19=9500, 200*2=400, 50*1=50, 20*2=40, 5*1=5, 2*2=4 = 9999? 9500+400=9900, +50=9950, +40=9990, +5=9995, +4=9999 ✓ = 19+2+1+2+1+2=27 coins
		if result != 27 {
			t.Errorf("coinChange(%v, 9999) = %d, expected 27", coins, result)
		}
	})

	t.Run("need max coins - only coin is 1", func(t *testing.T) {
		result := coinChange([]int{1}, 9999)
		if result != 9999 {
			t.Errorf("coinChange([1], 9999) = %d, expected 9999", result)
		}
	})

	t.Run("two coin types - one optimal", func(t *testing.T) {
		// 30 = 10+10+10 (3 coins) or 15+15 (2 coins)
		result := coinChange([]int{10, 15}, 30)
		if result != 2 {
			t.Errorf("coinChange([10,15], 30) = %d, expected 2", result)
		}
	})

	t.Run("impossible even with many denominations", func(t *testing.T) {
		// All coins are multiples of 3, amount is not
		result := coinChange([]int{3, 6, 9, 12}, 10)
		if result != -1 {
			t.Errorf("coinChange([3,6,9,12], 10) = %d, expected -1", result)
		}
	})
}

func BenchmarkCoinChange(b *testing.B) {
	coins := []int{1, 2, 5}
	amount := 11

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coinChange(coins, amount)
	}
}

func BenchmarkCoinChangeLargeAmount(b *testing.B) {
	coins := []int{1, 2, 5, 10, 25, 50}
	amount := 9999

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coinChange(coins, amount)
	}
}

func BenchmarkCoinChangeManyCoins(b *testing.B) {
	coins := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	amount := 10000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coinChange(coins, amount)
	}
}
