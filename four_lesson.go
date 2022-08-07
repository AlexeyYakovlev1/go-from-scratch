package main

import "fmt"

func run() {
	defer handlerPanic() // defer - always run in end
	printMessage("From printMessage")

	var nums = []int{1, 2, 3}
	nums[4] = 4
}

func printMessage(text string) {
	fmt.Printf("%s\n", text)
}

func handlerPanic() {
	if r := recover(); r != nil { // recover() likes a panic()
		fmt.Println(r)
	}

	fmt.Printf("%s\n", "handlerPanic success")
}
