package main

import (
	"errors"
	"fmt"
	"reflect"
)

func run() {
	var str string = "Hello, world!"
	// var num int32 = 3
	// var unum uint = 4 // only > 0
	// var weight float32 = 32.21
	// var isTrue bool = true
	var a byte = 168
	var b rune = 'a' // only in ''

	message := []byte("asd") // asd to bytes

	fmt.Println(reflect.TypeOf(str)) // string
	fmt.Println(message, b)          // [97 115 100] 97
	fmt.Printf("%c\n", a)            // ¨

	one()

	msg, err := enterTheClub(19)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(msg)
	fmt.Println(findMin(1, 2, 31, 32, 48, 0, 238, 4382))
}

func one() {
	a, b, c := 1, 2, 3
	a, _, c = 1, 5, 3

	fmt.Println(a, b, c)
}

func enterTheClub(age int) (string, error) {
	if age >= 18 {
		return "Continue", nil
	}

	return "Break", errors.New("YOU SO YOUNG")
}

func check(message string) string {
	switch message {
	case "test":
		return "You write test"
	case "test2":
		return "You write test2"
	default:
		return ""
	}
}

func findMin(numbers ...int) int {
	var lenNums int = len(numbers)

	if lenNums == 0 {
		return 0
	}

	var min int = numbers[0]

	for _, i := range numbers {
		if i < min {
			min = i
		}
	}

	return min
}
