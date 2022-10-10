package main

import "fmt"

func main() {
	var key string

	for {
		fmt.Println("------------家庭收支记账-------------")
		fmt.Println("            1.收支明细")
		fmt.Println("            2.登记收入")
		fmt.Println("            3.登记支出")
		fmt.Println("            4.退出软件")
		fmt.Println("请选择1-4：")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("---------")
		}
	}
}
