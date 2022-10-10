package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func copy(dstfilename string,srcfilename string)(written int64,err error)  {
	srcfile ,err := os.Open(srcfilename)
	if err != nil {
	    fmt.Printf("open is %v",err)
	}
	defer srcfile.Close()
	reader :=bufio.NewReader(srcfile)
	dstfile , err := os.OpenFile(dstfilename, os.O_WRONLY | os.O_CREATE,06666)
	if err != nil {
		fmt.Printf("open is %v",err)
	}
	write := bufio.NewWriter(dstfile)
	defer dstfile.Close()
	return io.Copy(write,reader)
}
func main()  {
	srcfile := "D:/ac.jpg"
	dstfile :="D:/abc.jpg"
	_,err :=copy(dstfile,srcfile)
	if err != nil {
	    fmt.Printf("拷贝错误，err=%v",err)
	}else {
	    fmt.Printf("拷贝成功")
	}
}
//将D盘当中的ac这个图片拷贝到D盘生成abc这个图片