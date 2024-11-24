package sync_mutex

import (
	"fmt"
	"sync"
)

func Task5() {

	var counter int
	var mutex sync.Mutex
	const numGoroutines = 10
	const increments = 10

	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {

				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()

	fmt.Printf("Final counter value: %d\n", counter)

}
