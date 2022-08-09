package main

import "fmt"

type Person interface {
	PersonWithIsAdult
	PersonWithIsEmploy
}

type Programmer struct {
	fullName string
	age      int
	job      string
	height   float32
}

type PersonWithIsAdult interface {
	isAdult() bool
}

type PersonWithIsEmploy interface {
	isEmploy() bool
}

type Teacher struct {
	fullName string
	age      int
	job      string
	height   float32
}

func (p Programmer) isAdult() bool {
	return p.age >= 18
}

func (t Teacher) isAdult() bool {
	return t.age >= 18
}

func (p Programmer) isEmploy() bool {
	return p.job == ""
}

func (t Teacher) isEmploy() bool {
	return t.job == ""
}

func run() {
	var frontEndDeveloper = Programmer{
		"Alexey Yakovlev",
		17,
		"front-end developer",
		170,
	}
	var mathTeacher = Teacher{
		"Jo Person",
		37,
		"math teacher",
		172,
	}

	checkIsAdultAndEmploy(frontEndDeveloper)
	checkIsAdultAndEmploy(mathTeacher)

	printInterface(frontEndDeveloper)
	printInterface(mathTeacher)

	printInterface("qwe")
	printInterface(84)
	printInterface(false)

	printMoreAndMore(2, 1, 3, 5, 3, 1, 2, 3, 7)

	typeSwitch(frontEndDeveloper)
}

func printMoreAndMore(num ...int) {
	fmt.Printf("%d", num)
}

func checkIsAdultAndEmploy(person Person) {
	if !person.isAdult() || !person.isEmploy() {
		fmt.Printf("%s\n", "You baby or you not have a job(")
		return
	}

	fmt.Printf("%s\n", "Welcome, old man, which have a job)")
}

func typeSwitch(i interface{}) {
	switch value := i.(type) {
	case int:
		fmt.Println("int", value)
	case string:
		fmt.Println("string", value)
	case bool:
		fmt.Println("bool", value)
	default:
		fmt.Println("unknown type", value)
	}

	str, ok := i.(string) // i.(type) - привидение интерфейса к определенному типу

	if !ok {
		fmt.Printf("%s\n", "Interface is not string")
	}

	fmt.Println(str)
}

func printInterface(i interface{}) {
	fmt.Printf("%+v\n", i)
}
