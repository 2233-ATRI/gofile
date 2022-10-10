package main

import "fmt"

func main() {
	var cites map[string]string
	cites = make(map[string]string)
	cites["no1"] = "北京"
	cites["n02"] = "上海"
	cites["no3"] = "西安"
	for i, v := range cites {
		fmt.Println(i, v)
	}
	stduent := make(map[string]map[string]string)
	stduent["stu01"] = make(map[string]string)
	stduent["stu01"]["name"] = "tom"
	stduent["stu01"]["sex"] = "nan"
	stduent["stu01"]["addness"] = "neijing "
	stduent["stu02"] = make(map[string]string)
	stduent["stu02"]["name"] = "jart"
	stduent["stu02"]["sex"] = "wu"
	stduent["stu02"]["addness"] = "xian "
	for v1, r := range stduent {
		fmt.Println("v1 = ", v1)
		for k, v := range r {
			fmt.Println("k = ", k, "v = ", v)
		}
	}
}
