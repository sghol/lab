package calculator

import "fmt"

func SayHello(name string) {
	fmt.Printf("Hello %s\n", name)
}

// simple calculator
func Add(x int, y int) int {
	return x + y
}

func Sub(x int, y int) int {
	return x - y
}

func Mul(x int, y int) int {
	return x * y
}
