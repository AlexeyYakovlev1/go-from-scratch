package main

import (
	"errors"
	"fmt"
)

// Structurs and they methods
type contact struct {
	firstName   string
	lastName    string
	phone       string
	email       string
	address     string
	dateOfBirth string
}

func (c contact) printInfo() {
	var fullName string = c.firstName + " " + c.lastName

	fmt.Printf("Fullname: %s;\n Phone: %s;\n Email: %s;\n Address: %s;\n dateOfBirth: %s;\n",
		fullName, c.phone, c.email, c.address, c.dateOfBirth)
}

func (c *contact) setName(name string) {
	c.firstName = name
}

// contants
const winePrice int = 1000

func run() {
	// variables
	const firstName string = "Alexey"
	const age int = 17

	var lastName string
	lastName = "Yakovlev"

	weight := 78.3

	personInfo(firstName, lastName, age, weight)

	// Pointers needed for changed values for variables (arguments) in function
	someName := "Alex"
	pSomeName := &someName // pointer (stored address of memory in which stored value variable someName)

	fmt.Println("Variable someName:", someName)     // Alex
	fmt.Println("Variable pSomeName:", pSomeName)   // 0xc00004e240
	fmt.Println("Variable *pSomeName:", *pSomeName) // Alex (dereferencing - get the value that is stored at the address pointed to by the pointer)

	*pSomeName = "Alex (changed)"
	fmt.Println("Variable *pSomeName which changed:", *pSomeName) // Alex (changed)

	// Demonstration
	var nameForDemonstration string = "Old name"
	fmt.Println(nameForDemonstration) // Old name
	setName(&nameForDemonstration)
	fmt.Println(nameForDemonstration) // New name

	// Empty pointers
	var emptyPointer *string // <nil>
	fmt.Println(emptyPointer)

	change, err := buyWine(18, 907)

	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Success:", change)
	}

	// arrays (static length)
	contactsList := [3]string{"Alex", "Other", "Quod"}
	contactsList[1] = "Change name"
	fmt.Println(len(contactsList)) // length array - 3
	for i := range contactsList {
		fmt.Printf("%d:%s\n", i+1, contactsList[i])
	}

	// or
	// for i, value := range contactsList {
	// 	fmt.Printf("%d:%s\n", i+1, value)
	// }
	// for i := 0; i <= len(contactsList); i++ {
	// 	fmt.Printf("%d:%s\n", i+1, value)
	// }

	// slices (dynamic length)
	newSlice := []int{1, 2, 3}
	newSlice = append(newSlice, 4, 5, 6)

	for i := 0; i < len(newSlice); i++ {
		fmt.Printf("%d. %d\n", i+1, newSlice[i])
	}

	// maps
	newMap := map[string]string{"Property1": "Value1", "Property2": "Value2"}
	fmt.Println(newMap["Property1"]) // Value1

	for property, value := range newMap {
		fmt.Println(property, value)
	}
	delete(newMap, "Property2")

	// Structurs and they methods
	c1 := contact{
		firstName:   "Alex",
		lastName:    "Pops",
		phone:       "89213123",
		email:       "alx@gmail.com",
		address:     "Zalupkino",
		dateOfBirth: "11 october",
	}

	fmt.Println(c1)           // {Alex Pops 89213123 alx@gmail.com Zalupkino 11 october}
	fmt.Println(c1.firstName) // Alex
	c1.setName("Changed name")
	c1.printInfo()
}

func personInfo(firstName string, lastName string, age int, weight float64) {
	var fullName string = firstName + " " + lastName

	// terms
	if isAdult(age) {
		fmt.Println("You adult")
	} else {
		fmt.Println("You baby")
	}

	fmt.Printf("Full name: %s\nAge: %d\nWeight: %.1f\n", fullName, age, weight) // printf - print formatted
}

func isAdult(age int) bool {
	return age >= 18
}

// pointers
func setName(name *string) {
	*name = "New name"
}

// use errors
func buyWine(age, moneyAmount int) (int, error) {
	if isAdult(age) && moneyAmount >= winePrice {
		return moneyAmount - winePrice, nil // return int and (how error) nil because error not need
	} else if !isAdult(age) {
		return moneyAmount, errors.New("Go to drink tea")
	} else {
		return moneyAmount, errors.New("Your amount so small")
	}
}
