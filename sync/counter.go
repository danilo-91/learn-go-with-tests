package counter

import (
	"sync"
)

type Counter struct {
	mu  sync.Mutex
	val int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val += 1
}

func (c *Counter) Value() int {
	return c.val
}
