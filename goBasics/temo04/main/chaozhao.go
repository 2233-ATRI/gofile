package main

import "fmt"

func erfenchazhao(arr *[6]int, leftindex int, rightindex int, findval int) {
	if leftindex > rightindex {
		fmt.Printf("no")
		return
	}
	middle := (leftindex + rightindex) / 2

	if (*arr)[middle] > findval {
		erfenchazhao(arr, leftindex, middle-1, findval)
	} else if (*arr)[middle] < findval {
		erfenchazhao(arr, middle+1, rightindex, findval)
	} else {
		fmt.Printf("zhaodale")
	}
}
func main() {
	//name := [5]string{"王大", "王二", "网三", "王四", "王五"}
	//var name1 = ""
	//fmt.Scanln(&name1)
	//for i := 0; i < 5; i++ {
	//	if name[i] == name1 {
	//		fmt.Println("yes")
	//		break
	//	}
	//}
	arr := [6]int{1, 9, 12, 15, 49, 89}
	erfenchazhao(&arr, 0, len(arr)-1, 89)
}
