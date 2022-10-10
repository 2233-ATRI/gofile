package main

import "fmt"

func main() {
	//var chan1 chan int//可读可写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 20 //只能写
	var chan3 <-chan int
	fmt.Println(chan3) //只能读
}
