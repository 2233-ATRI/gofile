package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//D盘当中的ac这个文件 D:/ac.txt
type charcount struct {
	chcount    int //英文的个数
	numcount   int //数字个数
	spaceCount int //空格个数
	otherCount int //其他字段的个数
}

func main() {
	filename := "D:/ac.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("open is err ", err)
		return
	}
	defer file.Close()
	var count charcount
	reader := bufio.NewReader(file)
	fmt.Printf("%v", reader)
	for {
		str, err := reader.ReadString('\n')
		if err != io.EOF {
			break
		}

		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.chcount++
			case v >= ' ' && v <= '\t':
				count.spaceCount++
			case v >= '0' && v <= '9':
				count.numcount++
			default:
				count.otherCount++
			}
		}
	}
	fmt.Printf("字符个数为%v，数字个数%v 空格个数%v 其他字符个数%v", count.chcount, count.spaceCount, count.numcount, count.otherCount)
}
