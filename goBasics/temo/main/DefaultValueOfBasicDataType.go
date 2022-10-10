package main

import (
	"fmt"
	"strconv"
)

//如果没有使用到某个包但不想删可以使用_对这个忽略掉
func main() {
	var a int
	var b bool
	var c float64
	var d float32
	var e string
	fmt.Println(a, "\n", b, "\n", c, "\n", d, "\n", e)
	//可以使用%v进行输出变量的值

	//不同的数据类型必须要进行显式转换，不会进行自动转换
	var i int = 100
	//将i转换成为float类型
	var n1 float64 = float64(i)
	fmt.Printf("%v\n", n1)
	fmt.Printf("%T\n", n1)
	fmt.Printf("%T\n", i)
	//被转换的是变量的 存储数据，变量本身没有进行改变
	//将大的数据类型转换成为小弟数据类型编译不会报错但会出现溢出处理
	var num1 int64 = 9999
	var num2 int8 = 12
	var num3 int8 = int8(num1) + num2
	fmt.Println(num3)

	//关于对String的转换
	var name1 int = 99
	var um2 float64 = 23.69
	var cd bool = true
	var mychar byte = 'h'
	var str string
	str = fmt.Sprintf("%d", name1)
	fmt.Printf("str type %T str %q\n", str, str)
	str = fmt.Sprintf("%f", um2)
	fmt.Printf("str type %T str %q\n", str, str)
	str = fmt.Sprintf("%t", cd)
	fmt.Printf("str type %T str %q\n", str, str)
	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type %T str %q\n", str, str)

	//既可以使用Sprintf这个函数进行对一个数据类型转换成为String
	//也可以通过strconv包中的函数进行所有的数据类型转换
	var str1 string = "true"
	var qw bool
	qw, _ = strconv.ParseBool(str1)
	fmt.Printf("%t", qw)
	//在五十行注意，该函数会返回两个值，第一个是布尔值第二个是报错，如果不想获得errir可以使用_进行忽略
	//如果string转换其他数据类型的不是数字进行转换，那么会默认转换成为0

}
