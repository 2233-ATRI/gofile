package main

import "fmt"

type Point struct {
	a int
	b int
}

func main() {
	var az interface{}
	var aq Point = Point{2, 3}
	az = aq

	var b Point
	b = az.(Point) //作为类型断言，表示判断az是否指向point类型的变量，如果是就转成point，并且赋值b变量，否则报错
	fmt.Println(b)

	var wqh interface{}
	var lyq float64 = 3.2
	wqh = lyq
	y, ok := wqh.(float32)
	if ok == true { //进行检测，对上面那个float32进行处理，改为float64即正确
		fmt.Println(y)
	} else {
		fmt.Println("error")
	}
	fmt.Println("继续执行")
}

//即接口类型和具体类型之间的相互转换，具体类型可以直接转换成为接口类型，接口类型需要其他类型的类型断言
