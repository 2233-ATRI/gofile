package main

import "fmt"

func main() {
	//for i := 1; i < 100; i++ {
	//	if i%2 == 0 {
	//		continue
	//	}
	//	fmt.Println(i)
	//}
	var c int
	var a int = 0
	var b int = 0
	for {
		fmt.Scanln(&c)
		if c > 0 {
			a++
			continue
		} else if c < 0 {
			b++
			continue
		} else {
			break
		}
	}
	fmt.Println(a)
	fmt.Println(b)
}
