package channels_basic_operations

import (
	"fmt"
	"time"
)

func sendvalue(ch chan int) {
	time.Sleep(500 * time.Millisecond)

	for i := 1; i <= 5; i++ {

		ch <- i

	}
	close(ch)

}

func Task4() {

	ch := make(chan int)

	go sendvalue(ch)

	for value := range ch {
		fmt.Println(value)
	}

}
