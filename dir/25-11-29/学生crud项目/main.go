package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	ID   int
	Name string
	Age  int
}

var students = make(map[int]Student)
var reader = bufio.NewReader(os.Stdin)

func main() {
	for {
		fmt.Println("------ 学生管理系统 ------")
		fmt.Println("1. 添加学生")
		fmt.Println("2. 删除学生")
		fmt.Println("3. 更新学生")
		fmt.Println("4. 查看所有学生")
		fmt.Println("5. 退出")
		fmt.Print("请输入操作编号：")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			addStudent()
		case "2":
			deleteStudent()
		case "3":
			updateStudent()
		case "4":
			listStudents()
		case "5":
			fmt.Println("退出系统，拜拜！👋")
			return
		default:
			fmt.Println("❗ 无效选项，请重新输入")
		}
	}
}

func addStudent() {
	fmt.Print("请输入学生 ID：")
	id := readInt()

	if _, exists := students[id]; exists {
		fmt.Println("❗ 学生 ID 已存在，无法重复添加")
		return
	}

	fmt.Print("请输入学生姓名：")
	name := readString()

	fmt.Print("请输入学生年龄：")
	age := readInt()

	students[id] = Student{
		ID:   id,
		Name: name,
		Age:  age,
	}

	fmt.Println("✔️ 添加成功！")
}

func deleteStudent() {
	fmt.Print("请输入要删除的学生 ID：")
	id := readInt()

	if _, exists := students[id]; !exists {
		fmt.Println("❗ 学生不存在")
		return
	}

	delete(students, id)
	fmt.Println("✔️ 删除成功！")
}

func updateStudent() {
	fmt.Print("请输入要更新的学生 ID：")
	id := readInt()

	student, exists := students[id]
	if !exists {
		fmt.Println("❗ 学生不存在")
		return
	}

	fmt.Printf("当前姓名：%s，输入新姓名（直接回车表示不改）：", student.Name)
	name := readStringOptional()
	if name != "" {
		student.Name = name
	}

	fmt.Printf("当前年龄：%d，输入新年龄（直接回车表示不改）：", student.Age)
	ageStr := readStringOptional()
	if ageStr != "" {
		age, _ := strconv.Atoi(ageStr)
		student.Age = age
	}

	students[id] = student
	fmt.Println("✔️ 更新成功！")
}

func listStudents() {
	fmt.Println("------ 学生列表 ------")
	if len(students) == 0 {
		fmt.Println("暂无学生数据")
		return
	}

	for _, s := range students {
		fmt.Printf("ID: %d | 姓名: %s | 年龄: %d\n", s.ID, s.Name, s.Age)
	}
}

func readInt() int {
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		val, err := strconv.Atoi(input)
		if err != nil {
			fmt.Print("❗ 输入无效，请输入数字：")
			continue
		}
		return val
	}
}

func readString() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readStringOptional() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
