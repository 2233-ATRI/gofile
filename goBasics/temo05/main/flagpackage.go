package main

import (
	"flag"
	"fmt"
)

//&user 就是接受用户命令行中输入的 -u后面的参数
//“u”，即为-u 指定的参数
//“” ，默认值
//“用户名，默认为空”说明
func main() {
	var user string
	var pwd string
	var host string
	var port int
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")
	flag.Parse()
	//进行转换操作，文件当中必须使用该方法
	fmt.Printf("user = %v ,pwd = %v ,host = %v , port = %v", user, pwd, host, port)
}
