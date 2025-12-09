package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func main() {
	UnmarshalMap()
}

func CMarshal() {
	p := Person{
		Age:  18,
		Name: "张三",
	}

	if str, err := json.Marshal(p); err != nil {
		fmt.Println("json序列化失败")
	} else {
		fmt.Println(string(str))
	}
}

func Unmarshal() {
	jsonStr := `{"name":"Bob","age":30}`

	var u Person
	if err := json.Unmarshal([]byte(jsonStr), &u); err != nil {
		panic(err)
	}

	fmt.Printf("Name: %s, Age: %d\n", u.Name, u.Age)
}

func UnmarshalMap() {
	// jsonStr := `[{"name":"A","age":20}, {"name":"B","age":22}]`
	jsonStr := `{"name":"A","age":20}`
	var m map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &m)

	fmt.Println(m["name"])
	fmt.Println(m["age"])
}
