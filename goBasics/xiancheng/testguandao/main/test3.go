package main

import (
	"fmt"
	"time"
)

func putnum(intchan chan int) {
	for i := 2; i < 80; i++ {
		intchan <- i
	}
	close(intchan)

}

func primnum(intchan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool

	for {

		num, ok := <-intchan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("有一个primrnum协程因为取不到数据推出")
	//	但不可以关闭这个管道，可以向exitchan
	exitChan <- true
}
func main() {
	intchan := make(chan int, 100)
	primechan := make(chan int, 2000) //放入结果
	//标记出来推出的管道
	exitchan := make(chan bool, 8)

	start := time.Now().Unix()
	go putnum(intchan)
	for i := 0; i < 8; i++ {
		go primnum(intchan, primechan, exitchan)
	}
	go func() {
		for i := 0; i < 88; i++ {
			<-exitchan
		}
		close(primechan)
	}()
	end := time.Now().Unix()
	fmt.Println(end - start)//计算用时

	for {
		res, ok := <-primechan
		if !ok {
			break
		}
		fmt.Println("素数", res)
	}
	fmt.Println("main线程推出")
}
