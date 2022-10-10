package main

import "fmt"

type Cat struct {
	Name  string
	Age   int
	Clour string
}

func main() {
	var cat1 Cat
	cat1.Name = "xiaobai"
	cat1.Age = 3
	cat1.Clour = "bai"
	fmt.Println(cat1)
}
