package ctx_pkg

import (
	"context"
	"fmt"
	"time"
)

func Task1() {

	ctx, cancel := context.WithCancel(context.Background())

	go printNumbers(ctx)

	time.Sleep(3 * time.Second)

	cancel()

	time.Sleep(1 * time.Second)

	fmt.Println("Main function done")

}

func printNumbers(ctx context.Context) {

	for i := 1; i <= 10; i++ {

		select {
		case <-ctx.Done():
			fmt.Println("Work Canceled", ctx.Err())
			return // remember

		default:
			fmt.Println(i)
			time.Sleep(1 * time.Second)

		}
	}
}
