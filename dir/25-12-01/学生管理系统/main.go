package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Id   int
	Name string
	Sex  int
}

var (
	students = map[int]Student{}
	reader   = bufio.NewReader(os.Stdin)
)

func main() {

	for {
		fmt.Println("请输入要进行的操作")
		fmt.Println("1. 添加学生")
		fmt.Println("2. 删除学生")
		fmt.Println("3. 修改学生")
		fmt.Println("4. 查询学生")
		fmt.Println("0. 退出")

		input := readInt()

		switch input {
		case 0:
			return
		case 1:
			add()
		case 2:
			del()
		case 3:
			edit()
		case 4:
			logStudents()

		default:
			fmt.Println("输入操作有误")
		}

		fmt.Println("")
	}
}

func add() {
	fmt.Println("请输入学生ID")
	id := readInt()
	fmt.Println("请输入学生姓名")
	name := readString()
	fmt.Println("请输入学生性别（1男 2女）")
	sex := readInt()

	stu := Student{
		Id: id, Name: name, Sex: sex,
	}
	students[id] = stu
	fmt.Println("新增成功")
}

func del() {
	fmt.Println("请输入学生ID")
	id := readInt()

	if val, exist := students[id]; exist {
		delete(students, id)
		fmt.Printf("学生 %s 已删除", val.Name)
	} else {
		fmt.Println("学生不存在")
	}
}

func edit() {
	fmt.Println("请输入要修改的学生ID")

	id := readInt()
	if _, exist := students[id]; exist {

	} else {
		fmt.Println("学生不存在")
		return
	}

	fmt.Println("请输入学生姓名")
	name := readString()
	fmt.Println("请输入学生性别（1男 2女）")
	sex := readInt()

	newStu := Student{
		Id: id, Name: name, Sex: sex,
	}
	students[id] = newStu
	fmt.Println("修改成功")
}

func logStudents() {
	for i, v := range students {
		var sexStr string
		if v.Sex == 1 {
			sexStr = "男"
		} else if v.Sex == 2 {
			sexStr = "女"
		} else {
			sexStr = "沃尔玛购物袋"
		}
		fmt.Printf("%d id: %d 名称 %s 性别: %s \n", i+1, v.Id, v.Name, sexStr)
	}
}

func readInt() int {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, _ := strconv.Atoi(input)
	return num
}

func readString() string {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}
