package main

import "fmt"

const (
	zero = iota
	one
	two
	three
)

func main() {
	fmt.Printf("zero: %v \n", zero)
	fmt.Printf("one: %v \n", one)
	fmt.Printf("two: %v \n", two)
	fmt.Printf("three: %v \n", three)
}
