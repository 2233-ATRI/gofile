//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	//var i int
//	//for i = 0; i < 10; i++ {
//	//	fmt.Println(i)
//	//}
//	////for a:=1;a <10; a++{
//	////    fmt.Println("hello word")
//	////}//另一种不同的循环初始值设置方式
//	//
//	//var az string = "hello word"
//	//for i := 0; i < len(az); i++ {
//	//	fmt.Printf("%c\n", az[i])
//	//} //面对字符串遍历传统方式
//	var aq string = "hello word，这个"
//	runes := []rune(aq)
//	for i, k := range aq {
//		fmt.Printf("%d %c\n", i, k)
//	} //Go语言特有的方式
//}

//package main
//
//import "fmt"
//
//func main() {
//	s := "阿福Chris"
//	_ = []rune(s)
//	fmt.Println(len(s))         //输出11
//	fmt.Println(len([]rune(s))) //输出7
//	for i := 0; i < len(s); i++ {
//		fmt.Printf("%c", s[i])
//	}
//}
package main

import "fmt"

func main() {
	var str string = "hello,word,北京"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c", str2[i])
	} //运行存在汉字的必须使用切片
}
