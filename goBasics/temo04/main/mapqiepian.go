package main

import "fmt"

func main() {
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "asd"
		monsters[0]["age"] = "500"
	}
	fmt.Println(monsters)
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "csd"
		monsters[1]["age"] = "500"
	}
	//if monsters[2] == nil {
	//	monsters[2] = make(map[string]string, 2)
	//	monsters[2]["name"] = "asd"
	//	monsters[2]["age"] = "500"
	//}
	//monsters = append(monsters, monsters[2])
	//fmt.Println(monsters)
	newmoster := map[string]string{
		"name": "sce",
		"age ": "200",
	}
	monsters = append(monsters, newmoster)
	fmt.Println(monsters)
}
