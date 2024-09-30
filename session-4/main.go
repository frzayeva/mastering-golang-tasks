package main

import "fmt"

// REFACTORING CODE INTO FUNCTIONS

// remaining calculations
func calculation_of_remaining_tasks(total, completed int) int {
	return total - completed
}

// task completion checks
func completion_task(total, completed int) bool {
	var isCompleted bool
	if total == completed {
		isCompleted = true
	} else {
		isCompleted = false
	}
	return isCompleted
}

// Define and use functions to print task statuses

func task_status(completed int) string {
	var status string
	if completed > 90 {
		status = "Completed "
	} else {
		status = "In progress"
	}
	return status

}

// Move task listing also to separate function
func task_listing(remaining int) {
	println("Task list:")
	for i := 1; i <= remaining; i++ {
		fmt.Printf("%d : Task %d\n", i, i)
	}

}

//Implement multiple return values

//Create a function that returns both the tasksRemaining and a boolean
//indicating if the project is near completion
//(e.g., if more than 90 tasks are completed)

func multiple_checking(total, completed int) (int, string) {

	remaining := calculation_of_remaining_tasks(total, completed)
	status := task_status(completed)

	return remaining, status
}

//3 Use recursion:

func task_countdown(remaining int) { //taskCount

	if remaining > 0 {
		fmt.Printf(" Tasks remaining : %d\n", remaining)
		task_countdown(remaining - 1)
	} else {
		fmt.Println(" All tasks completed!")
	}

}

/// change-it

// Variadic functions

//Implement a variadic function to accept and print the
//

func task_details_variadic(tasks ...string) {
	fmt.Println(" Task list:\n")

	for i, task := range tasks {

		fmt.Printf("%d : %s\n", i+1, task)
	}

}

func main() {

	total := 100
	completed := 95

	remaining, status := multiple_checking(total, completed)
	//remaining := calculation_of_remaining_tasks(total, completed)
	fmt.Printf("Task remaining : %d out of %d\n", remaining, total)

	fmt.Printf("Status check %s\n", status)
	//status := completion_task(total, completed)
	task_listing(remaining)
	fmt.Println("Recursive countdown of remaining tasks:")
	task_countdown(remaining)

	a := completion_task(total, completed)
	fmt.Println("Is the project completed? ", a)

	task_details_variadic("Task 1", "Task 2", "Task 3", "Task 4", "Task 5")

}
