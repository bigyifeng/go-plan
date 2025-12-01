package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	age  int
	name string
}

type SizeTest struct {
	age int
}

func main() {
	p1 := Person{
		age:  13,
		name: "hahah",
	}
	p2 := &Person{
		age:  111,
		name: "张三",
	}

	fmt.Println(p1.age)
	fmt.Println(p2.age)
	fmt.Println(unsafe.Sizeof(SizeTest{}))
}
