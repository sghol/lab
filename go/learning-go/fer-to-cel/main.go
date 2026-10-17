// Write a program that converts a temperature from Fahrenheit to Celsius.
// Use an integer variable fahrenheit set to 68.
// Print the result in Celsius as an integer.

// `Hint: Celsius = (Fahrenheit - 32) * 5 / 9. Use integer arithmetic.`

package main

import "fmt"

func main() {
	fer := 78
	c := (fer - 32) * 5 / 9
	fmt.Println("Fahrenheit:", fer, "Celsius:", c)
}
