package main

import (
	"fmt"
	"time"
)

func main() {
	//使用select 可以解决管道取数据的阻塞问题】
	intchan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intchan <- i
	}
	stringchan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringchan <- "hello" + fmt.Sprintf("%d", i)
	}
	//传统情况在遍历管道时，不关闭管道就会导致阻塞导致deadlock
	for {
		select {
		case v := <-intchan: //这里的intchan一直没有关闭，不会一直阻塞而会deadlock
			fmt.Printf("从intchan读取到数据%d", v)
			fmt.Println()
			time.Sleep(time.Second)
		//会向下自动配对到下一个case
		case v := <-stringchan:
			fmt.Printf("从stringchan读取到数据%s", v)
			fmt.Println()
			time.Sleep(time.Second)
		default:
			fmt.Printf("无法提取，可以加入逻辑")
			time.Sleep(time.Second)
			break //最好使用return
		}
	}
}
