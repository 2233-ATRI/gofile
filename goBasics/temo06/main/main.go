package main

import (
	"fmt"
	"gitfile/gofile/goBasics/temo06/bears"
)

func main() {
	var stu = bears.NewStudent("tom", 15.9)
	fmt.Println(*stu)
}
