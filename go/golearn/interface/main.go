package main

import (
	"fmt"
)


type Action interface {
	Walk(speed int) string
}


type Person struct {
	name string
	age int
}

type Rat struct {
	age int
}

func (p Person) Walk(speed int) string {
	if speed < 0 {
		fmt.Println("I don't move")
		return ""
	}
	fmt.Printf("I am walking with speed %d\n", speed)
	return ""
}

func (r Rat) Walk(speed int) string {
	if speed < 0 {
		fmt.Println("Rat is not moving.")
		return ""
	}
	if speed > 1 && speed < 5 {
		fmt.Println("Rat moves slowly")
		return ""
	}
	fmt.Printf("Rat moves fast with speed %d\n", speed)
	return ""
}


func main() {
	fmt.Println("Hello, World!")

	// peter := Person{name: "Peter Parker", age: 23}
	// peter.Walk(3)

	rat := Rat { age: 2 }
	rat.Walk(5)

}
