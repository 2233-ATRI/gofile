package main

import "fmt"

func main() {
	var name string
	var age int
	var sal float64
	var ispass bool
	fmt.Println("输入姓名")
	fmt.Scanln(&name)

	fmt.Println("输入年龄")
	fmt.Scanln(&age)
	//fmt.Scanf("%d",&age)
	fmt.Println("输入薪水")
	fmt.Scanln(&sal)
	fmt.Println("输入是否通过考试")
	fmt.Scanln(&ispass)
	fmt.Println(name, age, sal, ispass)
	
	var wq int = 011
	fmt.Println(wq)
}
