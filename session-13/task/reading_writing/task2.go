package reading_writing

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Task2() {

	file, err := os.Open("/home/farida/Documents/Go_HW/mastering-golang-tasks/session-13/task/reading_writing/story.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("Error occured:", err)
		return // why here we should use return bcs isn't a function
	}

	r := bufio.NewReader(file)

	lineCount := 0

	for {

		_, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error occured:", err)
			return
		}

		lineCount++
	}

	fmt.Println("Total number of lines :", lineCount)
}

// can't understand why it gives results 2 lines
