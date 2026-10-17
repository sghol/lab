package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// odds := filterOdds(nums)
	// fmt.Println(odds)
	// slicePractice()
	copyExample()

	s := "س"

	fmt.Println(s[0], s[1])

}

// isOdd return true if the number was odd number
func isOdd(n int) bool {
	return n%2 != 0
}

// filterOdds extract odd numbers and return them
func filterOdds(ns []int) []int {
	odds := []int{}

	for _, n := range ns {
		if isOdd(n) {
			odds = append(odds, n)
		}
	}
	return odds
}

func slicePractice() {
	x := make([]int, 0, 5)
	x = append(x, 3)

	fmt.Println(len(x), cap(x))
	fmt.Println("first index:", x[0])
}

func copyExample() {
	x := []int{1, 2, 3, 4, 5}
	y := make([]int, len(x))

	num := copy(y, x)

	fmt.Println(y, num)
}
