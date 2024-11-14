package sync_waitgroup

import (
	"fmt"
	"sync"
)

func calculateSum(nums []int, ch chan int) {

	sum := 0

	for _, num := range nums {
		sum += num
	}

	ch <- sum
}

func Task4() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	part1 := numbers[:len(numbers)/2]
	//fmt.Println(part1)

	part2 := numbers[len(numbers)/2:]

	sumChannel := make(chan int, 2)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		calculateSum(part1, sumChannel)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		calculateSum(part2, sumChannel)
	}()

	wg.Wait()
	close(sumChannel)

	totalSum := 0

	for partialSum := range sumChannel {

		fmt.Printf("Partial sum : %d\n", partialSum)
		totalSum += partialSum
	}

	fmt.Printf("Total sum : %d\n", totalSum)
}
