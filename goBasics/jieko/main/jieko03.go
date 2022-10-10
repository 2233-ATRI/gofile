package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//type stu struct {
//}
//type ainterfinece interface {
//	test1()
//}
//type binterfinece interface {
//	test2()
//}
//type cinterfinece interface {
//	ainterfinece
//	binterfinece
//	test3()
//}
//
//func (a Stu) test1() {
//
//}
//func (a Stu) test2() {
//
//}
//func (a Stu) test3() {
//
//}
//func main() {
//	var a Stu
//	var stu ainterfinece = a
//	stu.test1()
//}
type Hero struct {
	Name string
	Age  int
}

type Heroslice []Hero

func (h Heroslice) Len() int {
	return len(h)
}

func (h Heroslice) Less(i, j int) bool {
	return h[i].Age < h[j].Age
}

func (h Heroslice) Swap(i, j int) {
	temp := h[i]
	h[i] = h[j]
	h[j] = temp
}
func main() {
	var intslice = []int{0, -1, 90, 7, 10}
	sort.Ints(intslice)
	fmt.Println(intslice)
	var heroes Heroslice
	for i := 0; i < 10; i++ {
		hero := Hero{
			fmt.Sprintf("yingxiong%d", rand.Intn(100)),
			rand.Intn(99),
		}
		heroes = append(heroes, hero)
	}
	for _, v := range heroes {
		fmt.Println(v)
	}
	sort.Sort(heroes)
	fmt.Println("----------------")
	for _, v := range heroes {
		fmt.Println(v)
	}
}
