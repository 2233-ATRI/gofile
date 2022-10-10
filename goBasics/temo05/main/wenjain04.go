package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "D:/ac.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666) //写文件的时候使用权限
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}
	defer file.Close() //结束关闭
	str := "hello word\n"
	writer := bufio.NewWriter(file) //使用带缓存的*writer
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//文件是先写入缓存当中，所以需要调用flush方法才可以正常的写入磁盘当中，不使用会丢失数据
	//wrinter是自带缓存的
	writer.Flush()
}
