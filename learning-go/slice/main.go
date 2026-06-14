package main

import "fmt"


func main() {
    a1 := []int{1, 2, 3}
    a2 := []int{4, 5, 6}
    a3 := append(a1, a2...)

    fmt.Println(a3)



}
