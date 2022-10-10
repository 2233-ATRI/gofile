package main

import (
	"errors"
	"fmt"
)

//函数去读取一个文件，假如文件传入不正确，返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}
func test05() {
	err := readConf("config.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println("test05继续进行")
}

func test5() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res = ", res)
}
func main() {
	//test5()
	//fmt.Println("报错为")
	test05()
	fmt.Println("main()下的代码")
}
