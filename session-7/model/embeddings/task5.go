package embeddings

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

type Employee struct {
	person Person

	EmployeeID int
	Position   string
}

func Task5() {
	e := Employee{
		person: Person{
			FirstName: "Ali",
			LastName:  "Aliyev",
		},
		EmployeeID: 12345,
		Position:   "Software Engineer",
	}
	fmt.Println(e)
}

// Can't understand how to print struct line by line
