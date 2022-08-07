package main

import "fmt"

// === structurs
type User struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

type DataBase struct {
	m map[string]int
}

// ===

// === constructors
func NewDataBase() *DataBase {
	return &DataBase{
		m: make(map[string]int),
	}
}

func NewUser(name, sex string, weight, height, age int) User {
	return User{
		name:   name,
		age:    age,
		sex:    sex,
		weight: weight,
		height: height,
	}
}

// ===

type Any struct {
	name string
	age  int
}

type Any2 int

func (a Any2) printNum() {
	fmt.Printf("%d\n", a)
}

func run() {
	user1 := User{"Alexey", 17, "Male", 55, 170}
	user2 := NewUser("Petr", "Male", 83, 168, 23)

	var num int = 21
	Any2(num).printNum()

	var arr = [2]Any{{"Alex", 12}, {"Alexey", 15}}
	fmt.Println(arr)

	fmt.Println(user1.name)
	fmt.Printf("%+v\n", user2)

	user1.printUserName()
	user1.changeUserName("Maksim")
	user1.printUserName()
}

// === methods for structurs
func (u User) printUserName() {
	fmt.Println(u.name)
}

func (u *User) changeUserName(name string) {
	u.name = name
}

// ===
