package main

import "fmt"

func main() {
	str := "hello Word! 世界"
	runeStr := []rune(str)
	byteStr := []byte(str)

	fmt.Println(runeStr)
	fmt.Println(byteStr)
}
