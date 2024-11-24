package defining

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (c Circle) Area() float64 {

	return c.Radius * c.Radius * math.Pi
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func Task1() {

	var circle Shape
	var rectangle Shape

	circle = Circle{Radius: 10}
	rectangle = Rectangle{
		width:  15,
		height: 20,
	}

	fmt.Printf("Circle area :%.2f\n", circle.Area())
	fmt.Printf("Recatangle area :%.2f\n", rectangle.Area())

}
