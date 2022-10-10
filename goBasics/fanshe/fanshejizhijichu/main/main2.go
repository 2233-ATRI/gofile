package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Printf("%v\n", rVal.Kind())
	rVal.Elem().SetInt(20)
	//	无法实现直接对地址的更改，需要将其中的地址转化成为value
	//然后实现对于指针的类似操作
	/*
	   num：= 9
	   ptr *int = &num
	   num2 ：= *ptr //这一段与rVal.Elem()相似
	*/

}
func main() {
	var num int = 10
	reflect01(&num)
	fmt.Println(num)

}
