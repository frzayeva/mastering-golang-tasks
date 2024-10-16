package switching

import "fmt"

func checktype(i interface{}) {

	//Write a program that takes an interface{} and uses a type switch to determine whether the underlying value is an int, string, or bool.
	//Print a specific message for each type.
	//If the type is not recognized, print "Unknown type".

	//Output

	//Type is int: 100
	//Type is string: Hi, Gophers
	//Type is bool: true
	//Unknown type
	//

	switch v := i.(type) {
	case int:
		fmt.Printf("Type is int: %d\n", v)
	case string:
		fmt.Printf("Type is string: %s\n", v)

	case bool:
		fmt.Printf("Type is bool: %t\n", v)

	default:
		fmt.Printf("Unknown type")

	}

}

func Task6() {

	var i interface{}
	i = 100

	checktype(i)
	i = " Hi, Gophers"
	checktype(i)

	i = true
	checktype(i)
	i = 3.14
	checktype(i)
}
