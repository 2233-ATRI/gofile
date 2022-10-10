package main

import "fmt"

func main() {
	var a string = "tehainwe"
	fmt.Println(a)
	//var b string = "hello "
	//b[0]='a'
	//go中字符串是不可以改变的
	//可以使用反引号进行原生输出，这样可以保留下来换行和特殊符号，防止被攻击
	var c = `
func main() {
	var a string = "tehainwe"
	fmt.Println(a)
	//var b string = "hello "
	//b[0]='a'
	//go中字符串是不可以改变的
	//可以使用反引号进行原生输出，这样可以保留下来换行和特殊符号，防止被攻击`
	fmt.Println(c)
	//字符串可以使用+进行拼接
	var as string = "hello " + "word "
	as += "bearst"
	fmt.Println(as)
	//如果需要换行写的话，需要将加号放在上一行
}
