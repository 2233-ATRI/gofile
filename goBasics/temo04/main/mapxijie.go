package main

import "fmt"

type stu struct {
	name   string
	age    int
	Addnes string
}

func modifyUser(user map[string]map[string]string, name string) {
	if user[name] != nil {
		user[name]["pwd"] = "888888"
	} else {
		user[name] = make(map[string]string)
		user[name]["pwd"] = "888888"
		user[name]["nikename"] = name
	}
}
func maper(map1 map[int]int) {
	map1[10] = 900
}
func main() {
	////map1 := make(map[int]int)
	////map1[1] = 9
	////map1[7] = 12
	////map1[3] = 23
	////map1[10] = 100
	////map1[23] = 34
	////maper(map1)
	////fmt.Println(map1)
	//student := make(map[string]stu)
	//stu1 := stu{name: "nary", age: 15, Addnes: "中国"}
	//stu2 := stu{name: "tom", age: 18, Addnes: "美国"}
	//student["no1"] = stu1
	//student["no2"] = stu2
	//fmt.Println(student)
	users := make(map[string]map[string]string)

	modifyUser(users, "ton")
	modifyUser(users, "jary")
	fmt.Println(users)
}
