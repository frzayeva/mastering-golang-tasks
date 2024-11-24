package goroutines

import (
	"fmt"
	"time"
)

func Alphabet() {

	letters := []string{"A", "B", "C", "D", "E"}

	for _, letter := range letters {
		fmt.Println(letter)
		time.Sleep(200 * time.Millisecond)
	}

}

func Numbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(300 * time.Millisecond)
	}

}

func Task2() {

	go Alphabet()
	go Numbers()

	fmt.Println("Main finished")

	time.Sleep(2 * time.Second)

}
