package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("D:/cs.txt") //打开文件
	if err != nil {
		fmt.Println("open file err = ", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file) //这里是缓存的写法
	for {
		str, err := reader.ReadString('\n')
		if err != io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束")
}

//带缓冲的情况，一般来说会对大文件的读取更加快速
