package buffered_unbuffered_channels

import "fmt"

func Task6() {

	ch := make(chan string)
	ch <- "Hello"

	fmt.Println(ch)

}
