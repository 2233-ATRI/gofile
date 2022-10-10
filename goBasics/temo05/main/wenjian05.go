package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filemath := "D:/ac.txt"
	filemath2 := "D:/cs.txt"
	data, err := ioutil.ReadFile(filemath)
	if err != nil {
		fmt.Printf("err is %v", err)
	}
	err = ioutil.WriteFile(filemath2, data, 0666)
	if err != nil {
		fmt.Printf("err is %v", err)
	}
}//改变ac影响cs中的内容
