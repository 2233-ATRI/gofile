package main

import "fmt"

type Usb interface {
	Start() //这里可以加形参还有返回值
	Stop()
} //这里的usb就体现了多态
type phone struct {
	Name string
}

type Camera struct {
}

func (p phone) Start() {
	fmt.Println("手机开始工作")
}
func (p phone) Stop() {
	fmt.Println("手机停止工作")
}
func (c Camera) Start() {
	fmt.Println("相机开始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {
}

//编写一个working这个方法，接收usb接口类型变量
//只要是实现了Usb接口，所谓实现Usb接口即为实现Usb接口的声明的所有方法
func (c Computer) working(usb Usb) { //多态参数
	//通过Usb接口变量调用Start还有Stop
	usb.Start()
	usb.Stop()
	phone, ok := usb.(phone)
	if ok {
		phone.Stop()
	}
}

func main() {
	var c = Computer{}
	var p = phone{}
	var co = Camera{}
	c.working(p)
	c.working(co)

	var usb [3]Usb
	usb[0] = phone{"尼康"}
	usb[1] = phone{"vivo"}
	usb[2] = phone{"xiaomi"}
	fmt.Println(usb)
}

//多态数组
