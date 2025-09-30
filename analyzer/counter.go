package analyzer

import "sync"

type Counter[T comparable] struct {
	s sync.Mutex
	m map[T]int
}

func NewCounter[T comparable]() *Counter[T] {
	return &Counter[T]{
		m: make(map[T]int),
	}
}

func (c *Counter[T]) Increment(key T) {
	c.s.Lock()
	defer c.s.Unlock()

	c.m[key]++
}

func (c *Counter[T]) Get(key T) int {
	c.s.Lock()
	defer c.s.Unlock()

	return c.m[key]
}

func (c *Counter[T]) Len() int {
	c.s.Lock()
	defer c.s.Unlock()

	return len(c.m)
}

func (c *Counter[T]) ForEach(fn func(key T, count int)) {
	c.s.Lock()
	defer c.s.Unlock()

	for k, v := range c.m {
		fn(k, v)
	}
}
