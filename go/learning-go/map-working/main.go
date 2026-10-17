package main


import (
	"fmt"
)


type person struct {
	name string
	age int
}

func main() {
	// fmt.Println("hello world");
	//
	// daysOfWeeks := map[int]string {
	// 	0: "Sun",
	// 	1: "Sat",
	// 	2: "Mon",
	// 	3: "Tue",
	// 	4: "Wed",
	// 	5: "Thu",
	// 	6: "Fri",
	// }
	//
	// day := 3
	//
	// v, ok := daysOfWeeks[day]
	// if !ok {
	// 	fmt.Println("key not found")
	// }
	//
	// fmt.Println("The day is: ", v)

	var john person

	fmt.Println(john.age)


	bob := person{}

	fmt.Println(bob.age)

}

