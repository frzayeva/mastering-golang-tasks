package main

import "fmt"

func doubleVariadicElements(numbers ...*int) {
	for _, number := range numbers {

		*number = *number * 2
		//println(i)
		//println("------")
		//println(*number *2)
		//println("------")

	}

}

func main() {

	a, b, c := 4, 5, 6
	fmt.Printf("Before doubling : %d %d %d \n", a, b, c)
	doubleVariadicElements(&a, &b, &c)
	fmt.Printf("After doubling : %d %d %d ", a, b, c)

}
