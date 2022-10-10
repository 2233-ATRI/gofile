package main

import (
	"fmt"
	"sort"
)

func main() {

	map1 := make(map[int]int, 10)
	map1[0] = 100
	map1[4] = 29
	map1[8] = 24
	map1[9] = 28
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)
	for _, k := range keys {
		fmt.Println(map1[k])
	}
	//先把map的key顺序放进排序当中输出
	//	再对切片进行排序
	//遍历排序最后按照key输出map
}
