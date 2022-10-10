package main

import (
	"fmt"
	"net"
)

func main() {
	eonn, err := net.Dial("tcp", "192.168.20.253:8888")
	if err != nil {
		fmt.Println("err is = ", err)
		return
	}
	fmt.Println("conn 成功", eonn)

}
