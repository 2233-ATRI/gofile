package main

import "fmt"

func main() {
	str := "helloword&saw"
	slice := str[6:]
	fmt.Println(slice)
	//arr1 := []byte(str)
	arr1 := []rune(str)//可以修改中文
	arr1[0] = '被'
	arr2 := string(arr1)
	fmt.Println(arr2)
	//可以通过这个办法进行字符串的更改
}
