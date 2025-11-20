package main

import "fmt"

func main() {
	// log1To10()
	// sum1To100()
	// logEven1To50()
	// fn()
	fn2()
}

/* 1 */
func log1To10() {
	for i := range 10 {
		fmt.Println(i + 1)
	}
}

/* 2 */
func sum1To100() {
	sum := 0

	for i := range 100 {
		sum += i + 1
	}

	fmt.Println(sum)
}

/* 3 */
func logEven1To50() {
	for i := range 50 {
		if (i+1)%2 == 0 {
			fmt.Println(i + 1)
		}
	}
}

func fn() {
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
}

func fn2() {
	nums := []int{3, 7, 1, 9}
	for i, v := range nums {
		fmt.Println(i, v)
	}
}
