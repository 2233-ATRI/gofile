package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("D:/cs.txt") //打开文件
	if err != nil {
		fmt.Println("open file err = ", err)
	}
	fmt.Printf("file = %v", file)

	err = file.Close() //关闭文件
	if err != nil {
		fmt.Println("close file err ", err)
	}
}

//文件是一个指针类型
