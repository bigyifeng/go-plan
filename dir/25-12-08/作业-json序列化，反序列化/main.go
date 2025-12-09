package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	fn3()
}

func fn1() {
	p := Product{
		ID:    123,
		Name:  "张三",
		Price: 456456465,
	}

	data, _ := json.Marshal(p)
	fmt.Println(string(data))

}

func fn2() {
	str := `[
  {"id": 1, "name": "Pen", "price": 2.5},
  {"id": 2, "name": "Notebook", "price": 9.9}
]`
	var pArr []Product
	json.Unmarshal([]byte(str), &pArr)

}

func fn3() {
	if data, err := os.ReadFile("test.json"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(data))
		jsonStr := string(data)
		var obj map[string]interface{}
		json.Unmarshal([]byte(jsonStr), &obj)
		for i, v := range obj {
			fmt.Println(i, v)
		}
	}
}
