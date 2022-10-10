package main

import "fmt"

type Ainterface interface {
	Say()//接口中不可以存在常量
}
type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("stu")
}

type integer int

func (i integer) Say() {
	fmt.Println("intersay i = ", i)
}

type Binterrface interface {
	Hello()
}
type Monter struct {
}

func (m Monter) Hello() {
	fmt.Println("Monter Hello ()~~")
}
func (m Monter) Say() {
	fmt.Println("Monter Say()")

}
func main() {
	var stu Stu
	var a Ainterface = stu
	a.Say()

	var monster Monter
	var a2 Ainterface = monster
	var b2 Binterrface = monster
	a2.Say()
	b2.Hello()
}
