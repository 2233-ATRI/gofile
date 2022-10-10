package main

//序列化
import (
	"encoding/json"
	"fmt"
)

type monster struct {
	Name     string `json:"name"` //反射机制
	Age      int    `json:"age"`
	Birthday string
	sal      float64
	skill    string
}

func test() {
	master := monster{
		Name:     "wqn",
		Age:      231,
		Birthday: "2011.2.19",
		sal:      9.14,
		skill:    "sincq",
	}
	data, err := json.Marshal(master) //进行序列化操作
	if err != nil {
		fmt.Printf("序列号错误 err = %v", err)
	}
	fmt.Printf("序列化后为%v", string(data))
}

//对map进行序列化操作
func testmap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "bearst"
	a["age"] = 12
	a["addness"] = "sinafe"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("报错为%v", err)
	}
	fmt.Printf("a map 序列化后为%v", string(data))
}

//对切片进行序列化
func testlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["addness"] = "beijing"
	slice = append(slice, m1)
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err = %v", err)

	}
	fmt.Printf("a slice 序列化为%v", string(data))
}
func main() {
	//结构体,map，切片序列化
	test()
	fmt.Println()
	testmap()
	fmt.Println()
	testlice()
}
