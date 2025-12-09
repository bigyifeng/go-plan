package main

import (
	"fmt"
	"os"
)

func main() {
	GetFileInfo()
}

func GetFileInfo() {
	// 获取文件信息
	file, err := os.Stat("./dir/test1.txt")
	if os.IsNotExist(err) {
		fmt.Println("file is not exist")
	} else if err != nil {
		fmt.Println("文件访问错误")
	} else {
		fmt.Printf("文件 %s 访问成功", file.Name())
	}
}

func OpenFile() {
	// 查询文件，没有就新建
	file, err := os.OpenFile("./dir/test1.txt", os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()

	if os.IsNotExist(err) {
		fmt.Println("file is not exist")
	} else if err != nil {
		fmt.Println("文件访问错误")
	} else {
		fmt.Printf("文件 %s 访问成功", file.Name())
	}
}
