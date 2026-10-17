package main

import (
	"fmt"
	"github.com/ostadsgo/gogo/calculator"
)


// Person struct
type Person struct {
	Firstname string
	Lastname  string
	Age       int
	Email     string
}

func main() {
	// var age int
	// var name string
	//
	// name = "Peter Parker"
	// age = 45

	// calculator()

	// day := "Friday"
	// result1 := dayNumber(day)
	// fmt.Println(result1)
	// result2 := dayNumber2(day)
	// fmt.Println(result2)
	// result3 := dayNumber3(day)
	// fmt.Println(result3)
	// result4 := dayNumber4(day)
	// fmt.Println(result4)

	// strucExample()

	
	calc()





}

// functions example
func printMessage(name string, age int) {
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)
}

func calc() {
	a := 10
	b := 20
	fmt.Println("Simple Calculator")
	fmt.Println("-----------------")
	fmt.Printf("%d + %d = %d\n", a, b, calculator.Add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, calculator.Sub(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, calculator.Mul(a, b))
}


func dayNumber(day string) int {
	fmt.Println("--------------")
	fmt.Println("DAY NUMBER :: if else if")
	fmt.Println("--------------")

	if day == "Saturady" {
		return 0
	} else if day == "Sunday" {
		return 1
	} else if day == "Monday" {
		return 2
	} else if day == "Tuesday" {
		return 3
	} else if day == "Wednsday" {
		return 4
	} else if day == "Thursday" {
		return 5
	} else if day == "Friday" {
		return 6
	} else {
		return -1
	}

}

func dayNumber2(day string) int {
	fmt.Println("--------------")
	fmt.Println("DAY Number :: switch")
	fmt.Println("--------------")

	switch day {
	case "Sunday":
		return 0
	case "Saturady":
		return 1
	case "Monday":
		return 2
	case "Tuesday":
		return 3
	case "Wednsday":
		return 4
	case "Thursday":
		return 5
	case "Friday":
		return 6
	default:
		return -1
	}
}

func dayNumber3(day string) int {
	fmt.Println("--------------")
	fmt.Println("DAY NUMBER :: Slics")
	fmt.Println("--------------")

	days := []string{"Saturday", "Sunday", "Monday", "Tuesday", "Wednsday", "Thursday", "Friday"}
	for i, value := range days {
		if value == day {
			return i
		}
	}
	return -1
}

func dayNumber4(day string) int {
	fmt.Println("--------------")
	fmt.Println("DAY NUMBER :: Map")
	fmt.Println("--------------")

	days := make(map[string]int)

	days["Saturday"] = 0
	days["Sunday"] = 1
	days["Monday"] = 2
	days["Tuesday"] = 3
	days["Wednsday"] = 4
	days["Thursday"] = 5
	days["Friday"] = 6

	return days[day]
}

func strucExample() {
	peter := Person{
		Firstname: "Peter",
		Lastname:  "Parker",
		Age:       23,
	}

	fmt.Println(peter)
	peter.Info()
}

// method for struct
func (p *Person) Info() {
	fmt.Printf("First name: %s\n", p.Firstname)
	fmt.Printf("Last name: %s\n", p.Lastname)
	fmt.Printf("Age: %d\n", p.Age)
}
