package main

// Define interface
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}
