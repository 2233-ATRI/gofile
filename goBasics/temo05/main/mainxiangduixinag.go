package main

import "fmt"

type student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}
type box struct {
	long float64
	kuan float64
	gao  float64
}

func (a box) tiji() float64 {
	isq := a.long * a.gao * a.kuan
	return isq
}
func (student student) say() string {
	ifnww := fmt.Sprintf(student.name, student.age, student.id, student.score, student.gender)
	return ifnww
}

type menqiao struct {
	name string
	age  int
}

func (sc menqiao) menpiao() {

	if sc.age > 18 {
		fmt.Printf("%v的年龄是%v,门票价格是20", sc.name, sc.age)
	} else {
		fmt.Printf("%v的年龄是%v，门票免费", sc.name, sc.age)
	}

}
func main() {
	var student student = student{
		name:   "tom",
		gender: "male",
		age:    12,
		score:  15932,
	}
	var a = box{kuan: 12, long: 78, gao: 13}
	fmt.Println(a.tiji())
	fmt.Println(student)

	//fmt.Scanln(&)
	//var qw int
	//fmt.Scanln(&qw)
	//var sc = menqiao{name: as, age: qw}
	//fmt.Println(sc.menpiao)
}
