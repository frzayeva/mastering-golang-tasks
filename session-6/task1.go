package main

import "fmt"

//Task 1: Array Initialization and Basic Operations

//Create an integer array of length 5 and initialize it with values [10, 20, 30, 40, 50]

func sum(arr [5]int) int {
	sum := 0

	for _, value := range arr {

		sum += value
	}
	return sum
}

func reverse(arr [5]int) [5]int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr

}

func main() {
	A := [5]int{10, 20, 30, 40, 50}
	fmt.Printf(" First element : %d\n", A[0])
	fmt.Printf(" Last element : %d\n", A[len(A)-1])
	fmt.Printf("Sum : %d\n", sum(A))
	//fmt.Printf("Len : %d\n", len(A))
	reversedArr := reverse(A)
	fmt.Println(reversedArr)

}
