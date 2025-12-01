package main

import "fmt"

func main() {
	fmt.Println(*new(string))
	fmt.Println(*new(int))
	fmt.Println(*new([5]int))
	fmt.Println(*new([]float64))
}
