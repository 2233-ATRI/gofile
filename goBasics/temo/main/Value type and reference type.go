package main

import (
	"fmt"
)

//为了使用某个文件中的变量或者函数，我们需要先对这个包进行引入，具体我们可以使用这个包的相对路径
//基本数据类型int float 系列bool String数组还有结构体struct都是值类型，常在栈区
//引用类型包括指针，slice切片,map,管道chan ,interface都是属于引用类型，常在堆区，不使用便会作为垃圾进行处理
func main() {
	var num int = 10
	var Num int = 20
	fmt.Println(num, Num)
	//var numbers int =num++;不可以这样使用，++--都只能进行单独使用
	//if num++>0 {
	//
	//}这样也不可以使用，而且也不允许存在++i这样的运算符在前面的累计操作

}

//go中不存在public，pricate等对变量类型的情况
//如果首字母是大写是全局变量，如果首字母是小写则是局部变量
