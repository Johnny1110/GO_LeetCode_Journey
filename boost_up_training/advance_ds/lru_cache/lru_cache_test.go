package lru_cache

import (
	"testing"
)

func TestLRUCacheExample1(t *testing.T) {
	// LeetCode Example 1
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	if got := cache.Get(1); got != 1 {
		t.Errorf("Get(1) = %d, expected 1", got)
	}

	cache.Put(3, 3) // evicts key 2

	if got := cache.Get(2); got != -1 {
		t.Errorf("Get(2) = %d, expected -1 (should be evicted)", got)
	}

	cache.Put(4, 4) // evicts key 1

	if got := cache.Get(1); got != -1 {
		t.Errorf("Get(1) = %d, expected -1 (should be evicted)", got)
	}

	if got := cache.Get(3); got != 3 {
		t.Errorf("Get(3) = %d, expected 3", got)
	}

	if got := cache.Get(4); got != 4 {
		t.Errorf("Get(4) = %d, expected 4", got)
	}
}

func TestLRUCacheBasicOperations(t *testing.T) {
	t.Run("Single capacity cache", func(t *testing.T) {
		cache := Constructor(1)

		cache.Put(1, 10)
		if got := cache.Get(1); got != 10 {
			t.Errorf("Get(1) = %d, expected 10", got)
		}

		cache.Put(2, 20) // should evict key 1
		if got := cache.Get(1); got != -1 {
			t.Errorf("Get(1) = %d, expected -1 (evicted)", got)
		}
		if got := cache.Get(2); got != 20 {
			t.Errorf("Get(2) = %d, expected 20", got)
		}
	})

	t.Run("Get non-existent key", func(t *testing.T) {
		cache := Constructor(3)

		if got := cache.Get(1); got != -1 {
			t.Errorf("Get(1) on empty cache = %d, expected -1", got)
		}

		cache.Put(1, 10)
		if got := cache.Get(2); got != -1 {
			t.Errorf("Get(2) non-existent key = %d, expected -1", got)
		}
	})

	t.Run("Update existing key", func(t *testing.T) {
		cache := Constructor(2)

		cache.Put(1, 10)
		cache.Put(2, 20)
		cache.Put(1, 15) // update existing key

		if got := cache.Get(1); got != 15 {
			t.Errorf("Get(1) after update = %d, expected 15", got)
		}
		if got := cache.Get(2); got != 20 {
			t.Errorf("Get(2) = %d, expected 20", got)
		}
	})
}

func TestLRUCacheEvictionPolicy(t *testing.T) {
	t.Run("LRU eviction order", func(t *testing.T) {
		cache := Constructor(3)

		cache.Put(1, 10)
		cache.Put(2, 20)
		cache.Put(3, 30)

		// Access key 1 to make it most recently used
		cache.Get(1)

		// Add new item, should evict key 2 (least recently used)
		cache.Put(4, 40)

		if got := cache.Get(1); got != 10 {
			t.Errorf("Get(1) = %d, expected 10 (should not be evicted)", got)
		}
		if got := cache.Get(2); got != -1 {
			t.Errorf("Get(2) = %d, expected -1 (should be evicted)", got)
		}
		if got := cache.Get(3); got != 30 {
			t.Errorf("Get(3) = %d, expected 30", got)
		}
		if got := cache.Get(4); got != 40 {
			t.Errorf("Get(4) = %d, expected 40", got)
		}
	})

	t.Run("Put updates access order", func(t *testing.T) {
		cache := Constructor(2)

		cache.Put(1, 10)
		cache.Put(2, 20)
		cache.Put(1, 15) // update key 1, making it most recent
		cache.Put(3, 30) // should evict key 2, not key 1

		if got := cache.Get(1); got != 15 {
			t.Errorf("Get(1) = %d, expected 15", got)
		}
		if got := cache.Get(2); got != -1 {
			t.Errorf("Get(2) = %d, expected -1 (evicted)", got)
		}
		if got := cache.Get(3); got != 30 {
			t.Errorf("Get(3) = %d, expected 30", got)
		}
	})

	t.Run("Get updates access order", func(t *testing.T) {
		cache := Constructor(2)

		cache.Put(1, 10)
		cache.Put(2, 20)
		cache.Get(1)     // make key 1 most recent
		cache.Put(3, 30) // should evict key 2

		if got := cache.Get(1); got != 10 {
			t.Errorf("Get(1) = %d, expected 10", got)
		}
		if got := cache.Get(2); got != -1 {
			t.Errorf("Get(2) = %d, expected -1 (evicted)", got)
		}
	})
}

func TestLRUCacheEdgeCases(t *testing.T) {
	t.Run("Large capacity cache", func(t *testing.T) {
		cache := Constructor(1000)

		// Add many items
		for i := 1; i <= 500; i++ {
			cache.Put(i, i*10)
		}

		// Verify all items are accessible
		for i := 1; i <= 500; i++ {
			if got := cache.Get(i); got != i*10 {
				t.Errorf("Get(%d) = %d, expected %d", i, got, i*10)
			}
		}

		// Add more items to exceed capacity
		for i := 501; i <= 1200; i++ {
			cache.Put(i, i*10)
		}

		// First 200 items should be evicted
		for i := 1; i <= 200; i++ {
			if got := cache.Get(i); got != -1 {
				t.Errorf("Get(%d) = %d, expected -1 (should be evicted)", i, got)
			}
		}

		// Items 201-500 should still exist (accessed recently by Get calls)
		// Items 501-1200 should exist
		for i := 1001; i <= 1200; i++ {
			if got := cache.Get(i); got != i*10 {
				t.Errorf("Get(%d) = %d, expected %d", i, got, i*10)
			}
		}
	})

	t.Run("Negative keys and values", func(t *testing.T) {
		cache := Constructor(2)

		cache.Put(-1, -10)
		cache.Put(0, 0)

		if got := cache.Get(-1); got != -10 {
			t.Errorf("Get(-1) = %d, expected -10", got)
		}
		if got := cache.Get(0); got != 0 {
			t.Errorf("Get(0) = %d, expected 0", got)
		}
	})

	t.Run("Same key multiple updates", func(t *testing.T) {
		cache := Constructor(1)

		cache.Put(1, 10)
		cache.Put(1, 20)
		cache.Put(1, 30)

		if got := cache.Get(1); got != 30 {
			t.Errorf("Get(1) after multiple updates = %d, expected 30", got)
		}
	})
}

func TestLRUCacheComplexScenarios(t *testing.T) {
	t.Run("Alternating access pattern", func(t *testing.T) {
		cache := Constructor(3)

		cache.Put(1, 10)
		cache.Put(2, 20)
		cache.Put(3, 30)

		// Access pattern: 1, 2, 1, 2, 1, 2
		for i := 0; i < 3; i++ {
			cache.Get(1)
			cache.Get(2)
		}

		// Add new item, should evict key 3 (least recently used)
		cache.Put(4, 40)

		if got := cache.Get(3); got != -1 {
			t.Errorf("Get(3) = %d, expected -1 (should be evicted)", got)
		}
		if got := cache.Get(1); got != 10 {
			t.Errorf("Get(1) = %d, expected 10", got)
		}
		if got := cache.Get(2); got != 20 {
			t.Errorf("Get(2) = %d, expected 20", got)
		}
		if got := cache.Get(4); got != 40 {
			t.Errorf("Get(4) = %d, expected 40", got)
		}
	})

	t.Run("Mixed Put and Get operations", func(t *testing.T) {
		cache := Constructor(2)

		cache.Put(1, 1)
		cache.Put(2, 2)

		// Get 1 (makes 1 most recent)
		if got := cache.Get(1); got != 1 {
			t.Errorf("Get(1) = %d, expected 1", got)
		}

		// Put 3 (should evict 2)
		cache.Put(3, 3)

		// Get 2 (should be evicted)
		if got := cache.Get(2); got != -1 {
			t.Errorf("Get(2) = %d, expected -1", got)
		}

		// Put 4 (should evict 1)
		cache.Put(4, 4)

		// Get 1 (should be evicted)
		if got := cache.Get(1); got != -1 {
			t.Errorf("Get(1) = %d, expected -1", got)
		}

		// Get 3 and 4 (should still exist)
		if got := cache.Get(3); got != 3 {
			t.Errorf("Get(3) = %d, expected 3", got)
		}
		if got := cache.Get(4); got != 4 {
			t.Errorf("Get(4) = %d, expected 4", got)
		}
	})

	t.Run("Stress test with random operations", func(t *testing.T) {
		cache := Constructor(5)

		operations := []struct {
			op    string
			key   int
			value int
			want  int
		}{
			{"put", 1, 10, 0},
			{"put", 2, 20, 0},
			{"get", 1, 0, 10},
			{"put", 3, 30, 0},
			{"get", 2, 0, 20},
			{"put", 4, 40, 0},
			{"get", 1, 0, 10},
			{"get", 3, 0, 30},
			{"get", 4, 0, 40},
			{"put", 5, 50, 0},
			{"put", 6, 60, 0}, // should evict 2
			{"get", 2, 0, -1}, // should be evicted
			{"get", 1, 0, 10},
			{"get", 3, 0, 30},
			{"get", 4, 0, 40},
			{"get", 5, 0, 50},
			{"get", 6, 0, 60},
		}

		for i, op := range operations {
			switch op.op {
			case "put":
				cache.Put(op.key, op.value)
			case "get":
				if got := cache.Get(op.key); got != op.want {
					t.Errorf("Operation %d: Get(%d) = %d, expected %d", i, op.key, got, op.want)
				}
			}
		}
	})
}

func BenchmarkLRUCache(b *testing.B) {
	b.Run("Put operations", func(b *testing.B) {
		cache := Constructor(1000)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			cache.Put(i%1000, i)
		}
	})

	b.Run("Get operations", func(b *testing.B) {
		cache := Constructor(1000)

		// Pre-populate cache
		for i := 0; i < 1000; i++ {
			cache.Put(i, i*10)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cache.Get(i % 1000)
		}
	})

	b.Run("Mixed operations", func(b *testing.B) {
		cache := Constructor(100)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			if i%2 == 0 {
				cache.Put(i%100, i)
			} else {
				cache.Get(i % 100)
			}
		}
	})
}
