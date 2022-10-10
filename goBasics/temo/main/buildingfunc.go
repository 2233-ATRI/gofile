package main

import "fmt"

func main() {
	num1 := new(int)
	*num1 = 100
	fmt.Printf("类型 %T，值 %v, 地址 %v\n", num1, num1, &num1)
}
