package main

import "fmt"

//和c语言相似，同类型的指针只能存储同类型的数据地址 int的指针是*int float64的指针是*float64
//基本数据类型在内存中的布局
func main() {
	var i int = 10
	fmt.Println("i的地址是：", &i)
	var ptr *int = &i
	fmt.Println(ptr)
	fmt.Printf("ptr 的保存地址%v\n", ptr)
	fmt.Printf("ptr的地址%v\n", &ptr)
	fmt.Printf("ptr指向的值%v\n：", *ptr)
}
