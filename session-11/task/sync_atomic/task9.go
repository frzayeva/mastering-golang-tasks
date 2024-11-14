package sync_atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Task9() {

	var counter int32

	var wg sync.WaitGroup

	numGoroutines := 10
	increment := 10

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {

		go func() {

			defer wg.Done()
			for j := 0; j < increment; j++ {
				atomic.AddInt32(&counter, 1)
			}

		}()
	}

	wg.Wait()
	fmt.Printf("Final atomic value counter : %d\n", counter)

}
