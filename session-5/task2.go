package main

import "fmt"

func incrementbyValue(a int) int {
	a = a + 1
	return a
}

func incrementbyPointer(x *int) {
	*x = *x + 1
	// Question
	//why we can't use  ??
	//a =*x +1
	//return a

}

func main() {

	x := 5
	fmt.Printf("Value of x before incrementing by value: %d\n", x)
	incrementbyValue(x)

	fmt.Printf("Value of x after incrementing by value: %d\n", x)

	fmt.Printf("Value of x before incrementing by pointer: %d\n", x)

	incrementbyPointer(&x)
	fmt.Printf("Value of x after incrementing by pointer: %d\n", x)

}
