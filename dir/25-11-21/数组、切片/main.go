package main

import "fmt"

func main() {
	const length = 5
	nums := [length]int{1, 2, 3, 4, 5}

	qp := nums[1:4]
	fmt.Println(qp)

	qp[0] = 3
	fmt.Printf("%v\n", qp)
}
