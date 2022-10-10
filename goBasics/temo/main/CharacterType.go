package main

import (
	"fmt"
	"unsafe"
)

//go的字符串不是字符的处理而是字节串
func main() {
	var a byte = '9'
	var b byte = 'c'
	fmt.Printf("a=%c,b=%c", a, b)
	fmt.Println("a=", a, "b=", b)
	var c int = 22269
	fmt.Printf("%c", c)
	//可以直接通过%c进行对其他编码的处理，也可以进行码值运算
	var d = 10 + 'a'
	fmt.Printf("%c\n", d)
	var az bool = false
	fmt.Println("空间", unsafe.Sizeof(az))
}
