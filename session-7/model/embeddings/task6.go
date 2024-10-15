package embeddings

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

func Task6() {
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
