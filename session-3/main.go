package main

import "fmt"

func main() {

	// Operators
	totaltasks := 100
	completedtasks := 25
	remainingtasks := totaltasks - completedtasks

	fmt.Printf("Tasks remaining %d out of %d\n", remainingtasks, totaltasks)

	//	Conditional Statements
	var status string
	if remainingtasks == totaltasks {
		status = "COMPLETED"
	} else {
		status = "IN PROGRESS"
	}

	fmt.Printf("Current Project Status : %s\n", status)

	// switch
	var progress string

	switch {
	case completedtasks < 30:
		progress = "starting phase"

	case completedtasks >= 30 && completedtasks <= 60:
		progress = "midway"

	case completedtasks > 60:
		progress = "final phase"

	}

	fmt.Printf("Project is in the %s\n", progress)

	//LOOPS
	println(" Task list:")
	for i := 1; i <= remainingtasks; i++ {
		fmt.Printf(" -Task %d\n", i)

	}
}
