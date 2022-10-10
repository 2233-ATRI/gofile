package main

import "fmt"

//func (a A) test ()  {
//
//}
//表示有一个结构体,对于结构体A而言存在一个办法名为test
//（a A)表现test这个方法是和A类型绑定

type person struct {
	Name string
}

//func (p Person) speak() {
//	fmt.Println(p.Name, "a goodman")
//}
func (p person) test() {
	fmt.Println("test", p.Name)
}

//test和person绑定，只能通过person这个类型的变量进行调用，并可以通过其他方式进行调用
func main() {
	var p person
	p.Name = "tom"
	p.test()
}
