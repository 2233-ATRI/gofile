package main

import "fmt"

func writeData(intchan chan int) {
	for i := 0; i < 50; i++ {
		intchan <- i
		fmt.Println("写入的数据是", i)
	}
	close(intchan)
}
func readData(intchan chan int, exitchan chan bool) {
	for {
		v, ok := <-intchan
		if !ok {
			break
		}
		fmt.Println("读取到的数据是", v)
	}
	exitchan <- true
	close(exitchan)
}

func main() {
	intchan := make(chan int, 50)
	exitchan := make(chan bool, 1)
	go writeData(intchan)
	go readData(intchan, exitchan)
	for {
		_, ok := <-exitchan
		if !ok {
			break
		}
	}
}

//管道如果只有写没有读，那么会发生阻塞，但频率无所谓
