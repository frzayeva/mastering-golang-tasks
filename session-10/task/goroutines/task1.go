package goroutines

import (
	"fmt"
	"time"
)

func printnumbers() {

	for i := 1; i <= 5; i++ {
		println(i)
		time.Sleep(100 * time.Millisecond)
	}

}

func Task1() {
	fmt.Println(" Main function started")
	go printnumbers()

	time.Sleep(1 * time.Second)
	fmt.Println(" Main function ended")

}
