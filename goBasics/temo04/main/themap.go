package main

import "fmt"

func main() {
	//var a map[string]string
	//a = make(map[string]string, 10)//分配10个大小的空间
	////citise :=make(map[string]int)//也是一种map的方式
	//a["no1"] = "wqh"
	//a["no2"] = "qh"
	//fmt.Println(a)

	stduent := make(map[string]map[string]string)
	stduent["stu01"] = make(map[string]string)
	stduent["stu01"]["name"] = "tom"
	stduent["stu01"]["sex"] = "nan"
	stduent["stu01"]["addness"] = "neijing "
	fmt.Println(stduent["stu01"]["name"])
	fmt.Println(stduent["stu01"])
	delete(stduent["stu01"], "name") //删除某个具体的
	fmt.Println(stduent["stu01"])
	//stduent = make(map[string]map[string]string)
	//	这样处理之后会导致map原先的全部内容删除

	vr, i := stduent["stu01"]
	if i {
		fmt.Printf("有stu01key 值为%v", vr)
	} else {
		fmt.Printf("meiyou\n")
	}
}

//map是一个无序的结构
//map中的key不可以重复，value可以重复
