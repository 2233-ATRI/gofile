package main

import "fmt"

//使用var 变量名 chan 数据结构进行定义声明
//管道存在类型，int 类型只能使用int chan

func main() {
	//创建一个可以存放3个int类型的管道
	var intchan chan int
	intchan = make(chan int, 3) //为其中分配内存其中可以存放三个int类型的数据
	fmt.Println("intchan 的值是", intchan)
	fmt.Println("地址为", &intchan)
	intchan <- 10
	num := 211
	intchan <- num
	fmt.Println(len(intchan)) //内存
	fmt.Println(cap(intchan)) //容量
	fmt.Println()
	//读取数据
	var num2 int
	num2 = <-intchan
	fmt.Println(num2)
	fmt.Println(len(intchan))
	fmt.Println(cap(intchan))
	//	一旦取出就可以在往其中再放入数据

} //一定不可以让内存超过容量
//管道中遵守先进先出
//管道存放的是地址
