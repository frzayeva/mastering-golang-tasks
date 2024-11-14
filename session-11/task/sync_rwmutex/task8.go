package sync_rwmutex

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.RWMutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
	fmt.Printf("Write updated counter :%d\n", c.value)
}

func (c *Counter) Read() int {

	c.mu.RLock()

	defer c.mu.RUnlock()

	fmt.Printf("Reader accessed counter :%d\n", c.value)
	return c.value
}

func Task8() {
	counter := Counter{}
	var wg sync.WaitGroup

	//Reader
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {

			defer wg.Done()

			for j := 0; j < 3; j++ {

				counter.Read()
				time.Sleep(100 * time.Millisecond)
			}
		}(i)

	}

	//	writer

	wg.Add(1)

	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			time.Sleep(100 * time.Millisecond)
			counter.Increment()
		}
	}()

	wg.Wait()

}
