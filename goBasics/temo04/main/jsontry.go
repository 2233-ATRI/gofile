package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Num int
}
type B struct {
	Num int
}
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	var a A
	var b B
	a = A(b)
	fmt.Println(a, b)

	Monster := Monster{"wqh", 12, "wui"}
	jsonStr, _ := json.Marshal(Monster)
	fmt.Println(string(jsonStr))
}
//序列化具体使用场景：用于客户端和服务端之间的传递，使服务端表现更加清楚

