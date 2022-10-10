package main

import "fmt"

func main() {
	var a int32
	var b int32
	fmt.Scanln(&a, &b)
	if a+b >= 50 {
		fmt.Println("hello word!")
	}
}
