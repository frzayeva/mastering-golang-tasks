package ctx_pkg

import (
	"context"
	"fmt"
	"time"
)

func Task2() {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go runningTask(ctx)

	time.Sleep(6 * time.Second)

}

func runningTask(ctx context.Context) {
	defer customDefer()

	select {

	case <-ctx.Done(): // why here we 're using <- , shouldn't we use it only in channels
		fmt.Println("Task cancelled : Timeout occurred")

	case <-time.After(5 * time.Second):
		fmt.Println("Task completed after 5 sec-s")

	}
}

func customDefer() {

	fmt.Println("Exiting the task")
}
