package sync_rwmutex

import (
	"fmt"
	"sync"
)

func Task7() {

	userData := make(map[string]int)
	//var mu sync.RWMutex
	mu := sync.RWMutex{}

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {

		wg.Add(2)

		//	Reader
		go func() {

			mu.RLock()
			_ = userData
			mu.RUnlock()
			defer wg.Done()
		}()

		//	Writer

		go func() {
			mu.Lock()
			userData["Garay"] = 20
			userData["Ali"] = 25
			userData["Madina"] = 28
			mu.Unlock()

			defer wg.Done()
		}()

	}
	wg.Wait()

	mu.RLock()
	defer mu.RUnlock()
	fmt.Println("Final User Data :", userData)

}
