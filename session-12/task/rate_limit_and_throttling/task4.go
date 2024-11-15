package rate_limit_and_throttling

import (
	"fmt"
	"time"
)

var rateLimit = 2
var jobCount = 0

func Task4() {

	var resetTime = 1 * time.Second
	ticker := time.NewTicker(resetTime)
	defer ticker.Stop()

	jobChan := make(chan struct{}, rateLimit)

	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		if jobCount < rateLimit {

			jobChan <- struct{}{}
			jobCount++
		} else {
			fmt.Println("Rate limit reached")

		}

		select {
		case <-jobChan:
			job()

		case <-ticker.C:
			jobCount = 0
			fmt.Println("rate limit reset")
		}

	}

	fmt.Println("All jpbs completed")

}

func job() {

	fmt.Println("job started")
	time.Sleep(1 * time.Second)
}
