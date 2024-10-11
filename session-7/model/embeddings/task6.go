package main

import "fmt"

type Vehicle struct {
	Make  string
	Model string
	Year  int
}

type Car struct {
	Vehicle  Vehicle
	FuelType string
}

func main() {
	c := Car{
		Vehicle: Vehicle{
			Make:  "Kia",
			Model: "K5",
			Year:  2022,
		},
		FuelType: "Gasoline",
	}

	fmt.Println(c)

}
