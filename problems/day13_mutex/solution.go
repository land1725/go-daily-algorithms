package day13_mutex

import (
	"sync"
	"sync/atomic"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	//fmt.Printf("SafeCounter Inc: %d\n", c.count)
	c.count++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	//fmt.Printf("SafeCounter Value:%d\n", c.count)
	return c.count
}

type AtomicCounter struct {
	value int64
}

func (c *AtomicCounter) Inc() {
	atomic.AddInt64(&c.value, 1)
}

func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}
