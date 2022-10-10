package main

import (
	"fmt"
	"gitfile/gofile/goBasics/temo03/utils"
)

func main() {
	fmt.Println("main")
	fmt.Println("Age =", utils.Age, "Name =", utils.Name)
} //润过在main和其他的包中都有变量定义init函数，优先执行其他包中的然后再执行main包中的流程
