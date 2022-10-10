package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	allchan := make(chan interface{}, 3)
	allchan <- 10
	allchan <- "tom jack"
	cat := Cat{"小花猫", 13}
	allchan <- cat

	<-allchan
	<-allchan
	//如果需要获得管道当中的第三个元素，必须将其中前两个元素推出
	//
	newcat := <-allchan
	fmt.Printf("newcat =%T,newcat = %v\n", newcat, newcat)
	a := newcat.(Cat) //类型断言

	fmt.Printf("newcat.name = %v\n", a.Name)
	//	不可以直接使用fmt.printf("newCat.name=%v",newcat.name)
	close(allchan)
	allchan2 := make(chan interface{}, 3)
	for i := 0; i < 3; i++ {
		allchan2 <- i * 2
	}
	//for i := 0; i < 3; i++{
	//    allchan2
	//}遍历管道的时候不可以使用普通的for循环
	close(allchan2)
	for v := range allchan2 {
		fmt.Println("v =", v)
	} //可以在这里只给一个参数v不用再给下标，即可实现对其遍历，但要注意在这之前就要对其进行关闭
	//	先关闭再遍历就可以正常推出
}

//close之后依然可以进行读取但不再可以写入