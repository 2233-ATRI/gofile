package main

import "fmt"

func sum(a int, b int) int {
	//当执行到defer时，暂时不执行，会将defer后面的语句压入独立的栈当中（defer栈）
	//当函数执行结束之后，再从defer栈中按照先入后出的方式出栈执行
	//在defer语句放入栈中，也会将相关的值拷贝放入栈中
	//通常是在创建资源之后可以立刻使用defer确保在函数使用结束之后可以对内存释放减少浪费
	defer fmt.Println("ok1 n1 =", a)
	defer fmt.Println("ok2 n2 =", b)
	a++
	b++
	res := a + b
	fmt.Println("ok3 res = ", res)
	return res
}
func main() {
	res := sum(10, 20)
	fmt.Println(res)
}
