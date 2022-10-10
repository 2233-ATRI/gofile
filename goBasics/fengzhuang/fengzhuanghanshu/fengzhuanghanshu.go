package fengzhuanghanshu

import "fmt"

type yuangong struct {
	Name string
	age  int
	sal  float64
}

func Newyuangong(name string) *yuangong {
	return &yuangong{
		Name: name,
	}
}

func (a yuangong) SetAge(age int) {
	if age > 0 && age < 150 {
		a.age = age
	} else {
		fmt.Println("输入错误")
		a.age = 0
	}
}
func (a yuangong) GetAge() int {
	return a.age
}

func (a yuangong) SetSal(sal float64) {
	if sal > 3000 && sal < 30000 {
		a.sal = sal
	} else {
		fmt.Println("输入错误")
	}
}
func (a yuangong) GetSal() float64 {
	return a.sal
}
