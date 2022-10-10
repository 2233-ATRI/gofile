package main

import (
	"fmt"
	"reflect"
)

type student1 struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main() {
	//初始化
	stu1 := student1{
		Name:  "小马",
		Score: 99,
	}
	//返回类型
	t := reflect.TypeOf(stu1)
	fmt.Printf("name:%s kind:%v\n", t.Name(), t.Kind())
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}
	//通过字段名返回指定结构体的信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}
