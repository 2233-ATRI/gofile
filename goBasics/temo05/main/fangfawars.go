package main

import "fmt"

func (p Person) speak() {
	fmt.Println(p.Name, "a goodman")
}

type Person struct {
	Name string
}

func (p Person) try(n int) {
	res := 0
	for i := 0; i < n; i++ {
		res += i
	}
	fmt.Println(res)
	fmt.Println(p.Name, "jieguoshi", res)
}
func (p Person) xiangjia(a int, b int) int {
	return a + b
}

func main() {
	var p Person
	p.Name = "jory"
	p.speak()
	var n int
	fmt.Scanln(&n)
	p.try(n)
	p.xiangjia(10, 20)

}
