package main

import "fmt"

type Student struct {
	Name   string
	Grades []int
}

func (s Student) average() float64 {
	total := 0
	for _, grade := range s.Grades {
		total += grade
	}
	return float64(total) / float64(len(s.Grades))

}

func (s Student) Comparewith(s2 Student) string {
	avg1 := s.average()
	avg2 := s2.average()

	fmt.Printf(" Student 1 :%s, Average Grade : %.2f\n", s.Name, avg1)
	fmt.Printf("Student 2: %s, Average Grade :%.2f\n", s2.Name, avg2)

	if avg1 > avg2 {
		return fmt.Sprintf("%s has a higher average grade", s.Name)
	} else {
		return fmt.Sprintf("%s has a higher average grade", s2.Name)
	}

}

func main() {

	s1 := Student{

		Name:   "Farida",
		Grades: []int{3, 4, 5, 6},
	}

	s2 := Student{

		Name:   "Ayla",
		Grades: []int{4, 6, 2, 7},
	}

	a := s1.Comparewith(s2)

	fmt.Println(a)

}
