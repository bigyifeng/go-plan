package main

import "fmt"

func setToZero(a *int) {
	*a = 0
}

func main() {
	n := make([]int, 5)
	fmt.Printf("%T\n", &n)
}
