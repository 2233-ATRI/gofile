package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	//	lock是一个全局的互斥锁
	//sync是包：
	//mutex 互
	//互斥Mutex
	lock sync.Mutex
)

//函数进行计算n!,并且让这个结果放入myMap
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//在这里我们将res放入mymap当中
	//	加锁
	lock.Lock()
	myMap[n] = res
	//	解锁
	lock.Unlock()
}
func main() {
	for i := 0; i < 20; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 10)
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d] = %d\n", i, v)
	}
	lock.Unlock()
}

//加锁可以解决但是主线程很难估计所有的协程完成使用的时间
//使用加锁完成通信，同样不利于多个协程对全局变量的读写操作
//最好使用channel
