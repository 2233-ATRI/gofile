package main

import (
	"fmt"
	"gitfile/gofile/goBasics/temo02/utils"
	//util"gitfile/gofile/goBasics/temo02/utils"//这样在该文件中调用utils包的函数就需要使用util而不是utils
)

//可以在引入的包前面更改使用的名字，但更改以后应该使用更改后的名字，原先的名字不再适用
func main() {
	var n1 float64 = 4
	var n2 float64 = 3
	var operator byte = '*'
	result := utils.Cal(n1, n2, operator)
	fmt.Println("result = ", result)
}
