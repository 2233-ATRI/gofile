package main

import "fmt"

type money struct {
	Accountno string
	pwd string
	balane float64
}

func (ac *money)cunkuan(rmb float64 ,pwd string)  {//存款
	if pwd != ac.pwd {
	    fmt.Println("密码错误")
		return
	}
	if rmb<0{
		fmt.Println("金额错误")
		return
	}
	ac.balane+=rmb
	fmt.Println("存款成功")
}

func (ac *money)qukuan(rmb float64 ,pwd string)  {//存款
	if pwd != ac.pwd {
		fmt.Println("密码错误")
		return
	}
	if rmb<0|| ac.balane>rmb{
		fmt.Println("金额错误")
		return
	}
	ac.balane-=rmb
	fmt.Println("取款成功")
}
func (ac *money)chaxun(pwd string)  {
	fmt.Println(ac.balane)
}
func main() {
	var ac money = money{
		Accountno: "bearst",
		pwd : ("666666"),
		balane: 100,
	}
	for {
		fmt.Println("请输入你的需求")
		fmt.Println("1.存款，2取款，3查询，4退出")
		var a int
		fmt.Scanln(&a)
		if a == 1 {
		    var s string
			fmt.Scanln(&s)
			var m float64
			fmt.Scanln(&m)
			ac.qukuan(m,s)
		}else if a== 2 {
			var s string
			fmt.Scanln(&s)
			var m float64
			fmt.Scanln(&m)
			ac.cunkuan(m,s)
		} else if a == 3 {
			var s string
			fmt.Scanln(&s)
			ac.chaxun(s)
		}else {
			fmt.Println("输入错误，已退出")
			break
		}
	}
}