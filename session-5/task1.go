package main

import "fmt"

func main() {

	var x int = 10
	var p *int = &x
	//p := &x

	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %d\n", p)
	fmt.Printf("Value at pointer : %d\n", *p)

}
