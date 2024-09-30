package main

import "fmt"

//func main() {
//	str := "Hi, Gophers!"
//
//	for k, v := range str {
//		fmt.Printf("Key # %d of the value %v \n", k, v)
//	}
//
//	fmt.Println(str)
//}

func printTasks(tasks ...string) {
	// Check if any tasks are passed
	if len(tasks) == 0 {
		fmt.Println("No tasks provided.")
		return
	}

	fmt.Println("Task list:")
	for i := range tasks {

		fmt.Printf("%d\n", i)
		//fmt.Printf("%s\n", task)

		//fmt.Printf("%d: %s\n", i+1, task)
	}
}

func main() {
	// Call printTasks with any number of tasks
	printTasks("Task 1", "Task 2", "Task 3", "Task 4", "Task 5")

	// You can also call it with a different number of tasks or none
	// printTasks("Task A", "Task B")
	// printTasks()  // No tasks passed
}
