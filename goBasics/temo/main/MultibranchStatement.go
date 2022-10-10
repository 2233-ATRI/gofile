package main

import "fmt"

func test(c byte) byte {
	return c + 1
}

//函数应该写在import下面，因为import是引用包

func main() {
	var c byte = 10

	var achievement float64
	fmt.Scanln(&achievement)
	if achievement == 100 {
		fmt.Println("bmw")
	} else if achievement > 80 && achievement < 99 {
		fmt.Println("iphone")
	} else if achievement > 60 && achievement < 79 {
		fmt.Println("ipad")
	} else {
		fmt.Println(test(c))
	}
	//switch分支不需要再使用break即每一个cas里面都会自动给匹配一个break
	//switch当中的数据类型应该相同
	//var n1 int64 =20
	//var n2 int32 =20
	//switch n1 {
	//case n1:
	//	fmt.Println("m1")
	//case n2:
	//	//会在这里报错，'n1' (类型 'int32' 和 'int64' 不匹配) 的 switch 中的 case 'n2' 无效
	//}
	//
}
