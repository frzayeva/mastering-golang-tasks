package sync_mutex

import (
	"fmt"
	"sync"
)

func Task6() {

	grades := make(map[string]int)
	var mutex sync.Mutex
	var wg sync.WaitGroup

	students := []struct {
		name  string
		grade int
	}{
		{"Garay", 90},
		{"Ali", 85},
		{"Medina", 78},
		{"Garay", 95},
		{"Ali", 88},
	}

	for _, student := range students {
		wg.Add(1)
		go func(name string, grade int) {

			defer wg.Done()
			mutex.Lock()

			grades[name] = grade
			mutex.Unlock()

		}(student.name, student.grade)

	}
	wg.Wait()

	fmt.Printf(" Final Grades :% v\n", grades)
}
