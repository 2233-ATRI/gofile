package main

import "fmt"

//通常在Init函数当中完成初始化工作
//但在定义全局变量一般来说是最早被初始化的
func init() {
	fmt.Println("init")
}
func main() {
	fmt.Println("main")
}

//输出
//init
//main
