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
