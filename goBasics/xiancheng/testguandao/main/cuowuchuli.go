package main

import (
	"fmt"
	"time"
)

func sayhello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello word")
	}
}

func test1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test发生错误", err)
		}
	}()
	//使用recover进行错误的捕获，避免单一管道影响到了整个程序的崩溃
	//那个可能出现报错那个中使用recover
	var mymap map[int]string
	mymap[0] = "golang"

}
func main() {
	go sayhello()
	go test1()
	for i := 0; i < 10; i++ {
		fmt.Println("main() ok", i)
		time.Sleep(time.Second)
	}
}
