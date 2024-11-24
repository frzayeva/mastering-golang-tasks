package main

import "fmt"

func main() {

	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Original array  :", x) // it was modified based on capacity

	a := x[2:6]
	fmt.Println("Sub-slice : ", a)
	b := append(a, 10, 11, 12)
	fmt.Println("Appended slice :", b)

	fmt.Println("Original array after apppending  :", x) // it was modified based on capacity

}

// as i understand the reason behind why Original array after appending is : [0 1 2 3 4 5 10 11 12 9]
// is a is a slice ( so it reference to x ,and after appending it changes the values for x as well as
// it doesn't create copy but share same array in backup
