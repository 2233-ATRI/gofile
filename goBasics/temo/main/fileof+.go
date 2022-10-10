package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n1 int = 12
	var n2 int = 23
	fmt.Println(n1 + n2)
	//查看某一个变量的数据类型
	fmt.Printf("n1 = %T", n1)
	//查看某一个变量占用的字节数
	fmt.Printf("n1= %T", n1, "字节数 %d", unsafe.Sizeof(n1))
}
