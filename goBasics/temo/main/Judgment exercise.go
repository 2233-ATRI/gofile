package main

import "fmt"

func main() {
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	} else {
		fmt.Println("ok2")
	}
	a := 10
	b := 20
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)

}
