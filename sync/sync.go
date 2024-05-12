package sync

import "sync"

// a Mutex is a mutual exclusion lock, the zero value for a Mutex is an unlocked mutex
// Mutux allows us to add locks to our data
// WaitGroup is a means of waiting for goroutines to finish jobs

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
