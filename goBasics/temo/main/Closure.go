package main

import (
	"fmt"
	"strings"
)

func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	} //把内部本应该销毁的变量n进行了保留，保存在内部函数当中（栈内）
	//	闭包是类，函数是操作，变量n是字段，函数和它使用到n构成闭包
	//反复调用f函数时，因为n是初始化操作，所以每一次的调用就会进行叠加
	//重点是具体引用了哪些变量
	//理解为初始化之后不会再初始化，会进行保留
}
func makeSunffix(suffix string) func(string2 string) string {
	return func(name string) string {
		//如果name没有指定的后缀，则加上，否则返回原来的名字
		if !strings.HasPrefix(name, suffix) { //判断s是否有后缀字符串suffix。
			return name + suffix
		}
		return name
	}
}
func main() {
	f := AddUpper()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
	fmt.Println(f(3)) //16

	a := makeSunffix(".jpg")
	fmt.Println("处理之后文件名为", a("winter")) //
}
