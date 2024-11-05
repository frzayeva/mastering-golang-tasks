package channels_basic_operations

import (
	"fmt"
	"time"
)

func sendValue(ch chan int) {
	time.Sleep(500 * time.Millisecond)

	ch <- 42

}

func Task3() {
	ch := make(chan int)

	go sendValue(ch)

	value := <-ch

	fmt.Println("Received Value:", value)

}
