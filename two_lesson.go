package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type Square struct {
	sideLength float32
}

type Circle struct {
	radius float32
}

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func run() {
	var square = Square{8}

	printShapeArea(square)
}

func printShapeArea(shape Shape) {
	fmt.Println(shape.Area())
}
