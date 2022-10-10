package main

import "fmt"

/*
匿名结构体即可处理继承的问题，在处理的时候不需要重新定义，只需要嵌套一个student匿名结构体即可
形式如下
*/
//type good struct {
//	Name string
//	student string
//	Age int
//}
//type acc struct {
//	good
//	money float64
//	parch string
//}
type student struct {
	Name  string
	Age   int
	Score int
}

func (stu student) ShowInfo() {
	fmt.Println("学生姓名", stu.Name, "学生年龄", stu.Age, "学生成绩", stu.Score)
}
func (stu *student) Setscore(score int) {
	stu.Score = score
} //不加地址不会获得70

type toyil struct {
	student
}

func (p toyil) Setpoyil() {
	fmt.Println("小学生考试中，，，")
}

type daxuesheng struct {
	student
}

func (p daxuesheng) Setpoyil() {
	fmt.Println("大学生考试中")
}

type A struct {
	Name string
	Age int
}
type B struct {
    Name string
	Age int
}
type c struct {
	A//除非这个结构体当中是存在其他元素，假如只有其他结构体，必须使用匿名结构体
	B
	//如果这里没有Name string的话，必须在下面引用的时候规定好是A,B哪一个结构体当中的Name
}
func main() {
	toyil := &toyil{}
	toyil.student.Name = "tom"
	toyil.student.Age = 10
	toyil.Setpoyil()
	toyil.student.Setscore(70)
	toyil.student.ShowInfo()
	//可以简化成为
	toyil.ShowInfo()
	//当这样使用但假如往其中赋值的话，那么只能往tuyil中赋值,但假如使用toyil.student.Name = "tom"往
	//	其中赋值的话是往student中赋值，上一种有可能不会被使用

}
