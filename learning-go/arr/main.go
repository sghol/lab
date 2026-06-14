package main

import "fmt"

func main() {
    a := []float64{15, 19, 11}

    total := sum(a)

	ave := total / float64(len(a))

	fmt.Printf("Average is: %.2f\n", ave)
}


func sum(ns []float64) float64 {
	var sum float64
    for _, n := range ns {
		sum += n
	}
	return sum
}
