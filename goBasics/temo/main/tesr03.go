package main

import "fmt"

//可变参数的使用,可变参数具体指代的是... 这里可以确保里面的参数为任意多
func getSumAndSub1(n1 int, args ...int) int {
	var sub int = 0
	for i := 0; i < len(args); i++ {
		sub += args[i]
	}
	return sub
}
func main() {
	res4 := getSumAndSub1(10, 0, -1, 90, 10, 10)
	fmt.Println(res4)
}
