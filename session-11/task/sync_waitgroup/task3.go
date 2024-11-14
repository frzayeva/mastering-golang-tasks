package sync_waitgroup

import (
	"fmt"
	"sync"
	"time"
)

func Task3() {

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(id int) {

			fmt.Printf("Goroutine %d is starting \n", id)
			defer wg.Done()
			time.Sleep(1 * time.Second)

			fmt.Printf("Goroutine %d is finished\n", id)

		}(i)

	}

	wg.Wait()
	fmt.Println("All goroutines have been completed")

}
