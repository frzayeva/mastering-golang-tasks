package buffered_unbuffered_channels

import (
	"fmt"
)

func buff_channel(ch chan int) {

	for value := range ch {

		fmt.Println(value)
	}
}

func Task5() {
	ch := make(chan int, 3)

	ch <- 10
	ch <- 20
	ch <- 30

	go buff_channel(ch)

	close(ch)
	//  can't understand why output is not shown? :((
}
