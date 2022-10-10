package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器开始监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err =", err)
		return
	}
	defer listen.Close() //延时关闭

	for {
		com, err := listen.Accept() //等待客户端链接
		fmt.Println("等待客户端链接")
		if err != nil {
			fmt.Println("Accent err =", err)

		} else {
			fmt.Println("Accent suc con = ", com)
		}
		//准备一个协程为客户端服务
	}
	fmt.Println("listen suc =", listen)
}
