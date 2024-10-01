package main

import "fmt"

func main() {

	//for i := 1; i < 5; i++ {
	//
	//}

	// 10,      9, 8, 7, 6, 5

	// 1, 2, 3
	//for i := 10; i >= 5; i-- {
	//
	//	// 9
	//
	//	for j := 1; j <= 3; j++ {
	//
	//		fmt.Printf("%d + %d = %d \n", i, j, i+j)
	//
	//	}
	//
	//}

	for i := 1; i <= 5; i++ {

		fmt.Printf("Start: %d \n", i)

		for j := 1; j <= 5; j++ {
			if i == 3 {
				break
			}

			fmt.Printf("Secondary Loop num %d \n", j)

		}

		fmt.Printf("First Loop num %d \n", i)

	}

}
