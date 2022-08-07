package main

import (
	"errors"
	"fmt"
)

func run() {
	var users = map[string]int{
		"Alexey": 17,
		"Petr":   83,
	}
	users["Maks"] = 83

	var age, exists = users["Kostya"]

	delete(users, "Petr")

	fmt.Println(len(users))

	if exists {
		fmt.Println(age)
	} else {
		fmt.Println(errors.New("Kostya not exists in map!"))
	}

	for key, value := range users {
		fmt.Println(key, value)
	}
}
