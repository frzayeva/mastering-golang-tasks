package main

import "fmt"

// defer

func main() {

	defer func() {

		if a := recover(); a != nil {
			fmt.Println("Recovered from panic")
		}

	}()

	panicalMethod()

	fmt.Println("End of main")

}

func panicalMethod() {
	defer func() {
		fmt.Println("Last words...")
	}()
	fmt.Println("Start point of func")

	fmt.Println("End point of func")
	panic("Something went wrong")

}
