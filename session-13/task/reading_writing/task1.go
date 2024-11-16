package reading_writing

import (
	"encoding/csv"
	"fmt"
	"strconv"

	"log"

	//"log"
	"os"
)

func Task1() {
	//wd, _ := os.Getwd()
	//fmt.Println("Current Working Directory:", wd)
	file, err := os.Open("/home/farida/Documents/Go_HW/mastering-golang-tasks/session-13/task/reading_writing/data.csv")

	if err != nil {
		log.Fatal("Error opening file", err)

	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal("Error reading file", err)
	}

	var passingStudent []string
	for i, record := range records {
		//fmt.Println(i)
		//fmt.Println("------------")
		//fmt.Println(record)
		//fmt.Println("------------")

		if i == 0 {
			continue
		}

		grade, err := strconv.Atoi(record[2])

		if err != nil {
			log.Fatal(err)
		}

		if grade > 60 {
			passingStudent = append(passingStudent, record[0])
		}

	}

	fmt.Println("Students with passing grades:")

	for _, student := range passingStudent {

		fmt.Printf("- %s\n", student)
	}

}
