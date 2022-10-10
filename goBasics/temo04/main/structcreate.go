package main

import "fmt"

type Werson struct {
	Name string
	Age  int
}

func main() {
	p2 := Werson{}
	p2.Name = "wqh"
	fmt.Println(p2.Name)
	var p3 *Werson = new(Werson)
	(*p3).Name = "wqh"
	(*p3).Age = 12
	fmt.Println(*p3)

}
