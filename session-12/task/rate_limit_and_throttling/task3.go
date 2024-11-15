package rate_limit_and_throttling

import (
	"fmt"
	"time"
)

func Task3() {

	ticker := time.NewTicker(500 * time.Millisecond)

	defer ticker.Stop()

	for i := 0; i < 5; i++ {

		<-ticker.C //here same question why <-

		fmt.Println("Task executed")
	}

	fmt.Println("Ticker stopped ")

}
