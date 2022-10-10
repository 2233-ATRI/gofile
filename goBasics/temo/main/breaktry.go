package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//var count int = 0
	////for i = 0; i != n; i++ {
	////	if n != 99 {
	////		i++
	////	} else {
	////		fmt.Println(i)
	////		fmt.Println(n)
	////	}
	////}
	//for {
	//    rand.Seed(time.Now().Unix())
	//    n := rand.Intn(100) + 1
	//    count++
	//    if (n == 99) {
	//        break
	//    }
	//}
	//
	//fmt.Println("生成99，供使用了", count)
	var a int = 0
	n := rand.Intn(100)
	for {
		if a+n > 20 {
			break
		} else {
			a = a + n
		}
	}
	fmt.Println(a)
	fmt.Println(n)
}
