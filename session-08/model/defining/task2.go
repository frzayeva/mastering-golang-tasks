package defining

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}

type Person struct{}

func (d Dog) Speak() string {

	return "Woof!"
}

func (p Person) Speak() string {

	return "Hello!"
}

func speaking(s Speaker) { // burda biz interface uchun method yazmish sayiliriq? s Speaker receivable sayilir?
	fmt.Println(s.Speak())
}

func Task2() {
	var d Speaker
	var p Speaker

	d = Dog{}
	p = Person{}
	speaking(d)
	speaking(p)

	//fmt.Printf(" Dog says %s\n", speaking(d))  //why this case is not working

}
