package main

import "fmt"

// 要求：
// 删除指定位置的元素
// 返回新的切片
// 不使用循环截取两段手拼（可自由选择实现方式）
func remove(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {
	s := []int{10, 20, 30, 40, 50}
	fmt.Println(remove(s, 2))

	reverse(s)
	fmt.Println(s)
}

func reverse(s []int) {
	sCopy := make([]int, cap(s), cap(s))
	copy(sCopy, s)
	for i, v := range sCopy {
		s[cap(s)-1-i] = v
	}
}
