package main

import "fmt"

func main() {

	// OPERATORS
	totaltasks := 100
	completedtasks := 90
	remainingtasks := totaltasks - completedtasks

	var status string

	if completedtasks > 90 {
		status = " NEAR COMPLETION"

	} else {
		status = " IN PROGRESS"
	}

	var isCompleted bool
	if totaltasks == completedtasks {
		isCompleted = true
	} else {
		isCompleted = false
	}

	// CONDITIONAL STATEMENTS
	var progress string

	switch {
	case remainingtasks < 30:
		progress = " starting phase."

	case remainingtasks > 30 && remainingtasks <= 60:
		progress = " midway."

	case remainingtasks > 60:
		progress = " final phase."
	}

	fmt.Printf("Tasks remaining %d out of %d\n", remainingtasks, totaltasks)
	fmt.Printf("Current project status:%s\n", status)
	fmt.Printf("Project is in %s\n", progress)

	// LOOPS

	fmt.Printf("Task list:\n")
	for i := 1; i <= remainingtasks; i++ {
		fmt.Printf("-Task %d\n", i)
	}
	// Boolean
	fmt.Printf("Is the project completed? %v", isCompleted)

}
