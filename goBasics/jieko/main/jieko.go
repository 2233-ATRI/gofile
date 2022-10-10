package main

import "fmt"

type Usb interface {
	Start() //这里可以加形参还有返回值
	Stop()
}
type phone struct {
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
func (c Computer) working(usb Usb) {
	//通过Usb接口变量调用Start还有Stop
	usb.Start()
	usb.Stop()

}

func main() {
	var c = Computer{}
	var p = phone{}
	var co = Camera{}
	c.working(p)
	c.working(co)
}

//接口本身不可以创建实例，但可以指向
//接口中所有的方法都是没有方法体的，即都是没有实现的方法
