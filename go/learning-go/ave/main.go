package main

import "fmt"


func main() {
	var a float32
	var b float32
	var c float32

	a = 19.50
	b = 16.50
	c = 18.75

	var ave float32
	ave = (a + b + c) / 3

	fmt.Printf("Average is: %.2f\n", ave)
}
