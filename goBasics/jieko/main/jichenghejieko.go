package main

import "fmt"

type monkey struct {
	Name string
}
type Bird interface {
	flying()
}
type fish interface {
	swimming()
}

func (m monkey) climbing() {
	fmt.Println("会爬树")
} //这里是继承

type littlemonkey struct {
	monkey
} //继承

func (s littlemonkey) flying() {
	fmt.Println("通过学习学会飞行")
}
func (s littlemonkey) swimming() {
	fmt.Println("通过学习学会游泳")
}
func main() {
	monkey := littlemonkey{
		monkey{
			"bearst",
		},
	}
	monkey.climbing()
	monkey.swimming()
	//接口当一个结构体继承了另一个结构体，那么A结构体就自动继承了B结构体的字段和方法，而且可以直接使用
	//	如果新结构体不想破坏继承关系，又需要扩展功能，那么就可以使用接口，接口是继承的补充
}

//继承可以解决代码的复用性和可维护性
//接口更多是体现在设计上，对某个进行规范
//接口比地址更加灵活同时可以在一定程度上实现像什么的情况，继承只能实现是什么的情况
