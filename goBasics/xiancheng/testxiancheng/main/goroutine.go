package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test Hello word" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main() {
	go test() //go开启协程
	for i := 0; i < 10; i++ {
		fmt.Println("main Hello golang" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

//主线程结束协程会直接结束
//协程结束主线程运行不会受到什么其他的影响
