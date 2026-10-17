package main


import (
	"fmt"
)

// exercise 1.1
func hello_world() {
	fmt.Println("Hello, World!");
}



// exercise 1.1
func weight_on_mars() {
	fmt.Print("My weight on Mars is ")
	fmt.Print(68 * 0.3783)
	fmt.Print("KG, and I would be ")
	fmt.Print(37 * 365 / 687)
	fmt.Print(" Years old.\n")

}

func main() {
	weight_on_mars()
}
