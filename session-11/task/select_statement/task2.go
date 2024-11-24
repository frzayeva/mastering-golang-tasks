package select_statement

import (
	"fmt"
	"time"
)

func waitfunc(ch <-chan string) string {
	select {
	case s := <-ch:
		return s

	case <-time.After(3 * time.Second):
		return "timeout"

	}

}

func Task2() {

	ch := make(chan string)

	result := waitfunc(ch)
	fmt.Println(result)
}
