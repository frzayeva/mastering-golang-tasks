//Task 5: Type Assertions
//
//Write a program where you define an empty interface any and assign values of different types (e.g., int, string, float64) to it.
//Use type assertions to detect the underlying types and print different messages depending on the type.

package assertion

import "fmt"

func type_checking(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Value is of type int:%d\n ", v)

	case float64:
		fmt.Printf("Value is of type float64:%f\n ", v)

	case string:
		fmt.Printf("Value is of type string:%s\n ", v)

	}

}

func Task5() {
	var i any

	i = 42
	type_checking(i)

	i = "Golang"
	type_checking(i)

	i = 3.145
	type_checking(i)

}
