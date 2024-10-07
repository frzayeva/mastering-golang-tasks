package main

import "fmt"

//Create an empty slice using the make() function with an initial length of 3 and capacity of 5.
//Append values to the slice until its length exceeds 5.
//At each step, print the length and capacity of the slice to show how Go manages memory allocation internally.

func main() {

	x := make([]int, 3, 5)
	//x := make([]int, 0

	fmt.Printf("Initial length : %d, capacity :%d\n", len(x), cap(x))

	//x = append(x, 1)

	for i := 0; i < 10; i++ {
		x = append(x, i)
		fmt.Printf("After appending %d element, the len is %d  and cap is %d\n", i, len(x), cap(x))
	}

}
