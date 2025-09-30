package analyzer

import (
	"sync"
	"testing"
)

func TestIncrement(t *testing.T) {
	c := NewCounter[int]()

	// define total keys to add
	totalTestKeys := 100

	// create counter
	var wg sync.WaitGroup
	for i := 0; i < totalTestKeys; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			c.Increment(key)
		}(i)
	}
	wg.Wait()

	// tests all count values should be 1
	var keysCount int
	c.ForEach(func(key int, count int) {
		keysCount++

		if count != 1 {
			t.Errorf("expected 1 for key: %v, got %d", key, count)
		}
	})

	// test total keys
	if keysCount != totalTestKeys {
		t.Errorf("expected %v keys, got: %v", totalTestKeys, keysCount)
	}
}

func TestLen(t *testing.T) {
	c := NewCounter[int]()

	c.Increment(1)
	if got := c.Len(); got != 1 {
		t.Errorf("expected 1, got %d", got)
	}
}
