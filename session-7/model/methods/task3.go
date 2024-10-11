package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {

	return math.Pi * c.radius * c.radius

}

func main() {

	c := Circle{

		radius: 7.0,
	}

	fmt.Println("Circle Radius:", c.radius)

	fmt.Printf("Area :%.3f", c.area())

}
