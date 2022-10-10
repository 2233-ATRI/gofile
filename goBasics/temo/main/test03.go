package main

import "fmt"

//func test(n1 *int) {
//	*n1 = *n1 + 1
//	fmt.Printf("test() n1 =%d ", *n1)
//}//这个情况下会改变main栈中的n1
func test1(n1 int) {
	n1 = n1 + 1
	fmt.Println("test() n1 =", n1)
} //这种不会影响到main栈中的情况
func getnum(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}
func getSumAndSub(n1 int, n2 int) (int, int) {
	sun := n1 + n2
	sum := n1 - n2
	return sum, sun
}
func feibonaqieshulei(n1 int) int {

	if n1 == 1 || n1 == 2 {
		return 1
	} else {
		return feibonaqieshulei(n1-1) + feibonaqieshulei(n1-2)
	}
}
func main() {
	n1 := 10
	var n2 int = 12
	test1(n1)
	fmt.Println("main() n1 = ", n1)
	getnum(n1, n2)
	res1, res2 := getSumAndSub(n1, n2) //在这里可以使用下划线进行占位
	//_,res2 = getSumAndSub(n1,n2)//这样加法运算出来的结果就不会出现
	fmt.Println(res1, res2)
	a := getnum
	fmt.Printf("a的类型是%T,getSun的类型是%T", a, getSun) //a的类型是func(int, int) int,
	// getSun的类型是func(int, int) int

}

//在go这当中，函数可以当作一种数据类型，可以赋值给一个变量
//，则这个变量就是那个函数类型的变量通过该变量可以对函数调用
//go也支持函数作为形参进行调用
//可以使用type myint int 进行定义，实现都是int类型，但注意go依然认为这两个类型是两种类型

func getSun(n1 int, n2 int) int {
	return n1 + n2
}
