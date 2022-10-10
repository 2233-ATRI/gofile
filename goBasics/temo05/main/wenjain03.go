package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "D:/ac.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file %v", err)
	}
	//fmt.Printf("%v", content)//这样直接输出得到的不一样是[]byte的切片类型，所以显示的假如是汉字则会输出数字
	fmt.Printf("%v", string(content))
}

//直接将读取到底内容直接输出到终端里面
//这里不需要显示文件的open函数，所以在这里也不需要使用显示的close文件
//因为文件的close和open封装到了readfile内部
