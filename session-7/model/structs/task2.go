package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height

}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {

	//var rect Rectangle
	//
	//rect.width = 10.5
	//rect.height = 5.0

	rect := Rectangle{

		width:  10.5,
		height: 5.0, // , is importantt!
	}

	fmt.Println("Width:", rect.width)
	fmt.Println("Height:", rect.height)
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Area: %.2f", rect.Perimeter())

}
