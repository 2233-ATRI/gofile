package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Scares [5]float64
	pro    *int
	slice  []int
	map1   map[string]string
}

func main() {
	var p1 Person
	fmt.Println(p1)
	p1.slice = make([]int, 10)
	p1.slice[1] = 100
	fmt.Println(p1.slice)
	p1.map1 = make(map[string]string)
	p1.map1["name"] = "wqh"

	fmt.Println(p1.map1)
}
