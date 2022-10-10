package main

import "fmt"

var (
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
) //全局匿名函数
func main() {
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20) //匿名函数的使用办法
	fmt.Println(res1)
	//a := func(n1 int ,n2 int ) int  {
	//	return n1-n2
	//}
	//res2 := a(19,20)
	//fmt.Println(res2)//匿名函数的使用办法2
	//全局匿名函数的使用
	res4 := Fun1(4, 9)
	fmt.Println(res4)
}
