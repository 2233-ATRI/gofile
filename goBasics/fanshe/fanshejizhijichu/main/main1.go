package main

//使用反射获得变量的值，必须要求其数据类型匹配，比如说x是int，那么需要使用reflect.Value(x).Int()
//使用其他的可能会导致报错

//类型和类别有可能相同也有可能不相同
//对于普通的数据类型来说是相同的，但对于一个结构体来说
//类型是包名.student ，类别是struct
import (
	"fmt"
	"reflect"
)

//反射可以在运行的过程当中动态的的获得变量的各种信息，比如说类型类别等
//如果是结构体变量，还可以获得结构体本身的消息，包括结构体的字段和方法
//可以进行修改变量，调用关联关系
//
//type stu struct {
//
//}
//func test(b interface{})  {
//	var student stu
//	rVal := reflect.ValueOf(b)
//	//将接口interface转换成为reflect.Value
//	iVal := rVal.Interface()
//	//将reflect.Value转换成为接口
//	v:=iVal.(stu)
////将inteface 转换成为原来的变量类型
//}
func reflecttest(b interface{}) {
	//通过反射获得传入的变量的type,kind,值
	//	先获得reflect.type
	rtyp := reflect.TypeOf(b)

	fmt.Println(rtyp)
	fmt.Printf("真实类型%T\n", rtyp)
	//获取类型

	rVal := reflect.ValueOf(b)
	//其不是存粹的int类型，其无法直接被加减，真正的其实其类型是reflect.TypeOf，但也可以使用
	//以下办法进行更改
	//n2 := 2 +rVal//错误
	n2 := 2 + rVal.Int()
	fmt.Println(n2)
	fmt.Println(rVal)

	//	重新将rVal转换成为interface{}
	iv := rVal.Interface()
	num2 := iv.(int)
	fmt.Println("num2 = ", num2)
	fmt.Println()
	t := reflect.TypeOf(b)
	fmt.Println(t)
	v := reflect.ValueOf(b)
	k := v.Kind()
	fmt.Println(k)
	fmt.Println(v.Int())

}

type student struct {
	Name string
	Age  int
}

func reflectttest02(b interface{}) {

	rVal := reflect.ValueOf(b)

	//	重新将rVal转换成为interface{}
	iv := rVal.Interface()
	fmt.Printf("iv = %v\n  ivtype = %T\n", iv, iv)
	stu, ok := iv.(student)
	if ok {
		fmt.Printf("stu.name = %v\n", stu.Name)
	}
}
func main() {
	//var num int = 100
	//reflecttest(num)
	stu := student{
		Name: "tom",
		Age:  123,
	}
	reflectttest02(stu)
}
