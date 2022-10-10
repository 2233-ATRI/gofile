package main

import (
	"fmt"
	"math/rand"
	"time"
)

//import (
//	"fmt"
//)
//
//func main() {
//	var arr [6][9]int
//	arr[2][3] = 1
//	for i := 0; i < 6; i++ {
//		for j := 0; j < 9; j++ {
//			fmt.Print(arr[i][j], "")
//		}
//		fmt.Println()
//	}
//	for i, v := range arr {
//		for j, v2 := range v {
//			fmt.Printf("arr[%v][%v]=%v \t", i, j, v2)
//		}
//		fmt.Println()
//	}
//}
func main() {
	var arr [10]int
	var c int = 0
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		arr[i] = rand.Intn(100)
		c = c + arr[i]
	}
	c = c / 10
	for i := 0; i < 10; i++ {
		fmt.Println(arr[10-i])
	}
	var max = arr[1]
	for i := 0; i < 10; i++ {
		if max < arr[i] {
			arr[i] = max
		}
	}
	fmt.Println(max)
	for i := 0; i < 10; i++ {
		if arr[i] == max {
			fmt.Println(i)
		}
		if arr[i] == 55 {
			fmt.Println(i)
		}
	}
}
