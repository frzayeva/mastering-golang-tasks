package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	// Type Conversion

	// Output Values

	intE := math.MaxInt64
	floatE := math.MaxFloat64
	stringE := "44"

	res1 := int(floatE)
	res2 := float64(intE)
	res3 := strconv.Itoa(intE)
	res4, _ := strconv.Atoi(stringE)

	fmt.Println(floatE)

	fmt.Println(res1, res2, res3, res4)

	//str := 'Hello'

}
