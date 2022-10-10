package main

import "fmt"

func main() {
	//var az [7]int
	//for i := 0; i < 7; i++ {
	//	fmt.Scanln(&az[i])
	//}
	//var numArr01 [3]int =[3]int{1,2,3}
	//var numArr02  = [3]int {1, 2, 3}
	//var numArr03  = [...]int {1, 2, 3}//三个点是固定写法不可以更改
	//var numArr02  = [3]string {1:"tom", 0:"jack", 2:"marry"}
	//var achievement [5]float64
	//for i := 0; i < 5; i++ {
	//	fmt.Scanln(&achievement[i])
	//}
	//for i := 0; i < 5; i++ {
	//	fmt.Println(achievement[i])
	//	fmt.Printf("\n")
	//}
	//var zhimu [26]byte
	//zhimu[0] = 'A'
	//for i := 0; i < 26; i++ {
	//	zhimu[i] = zhimu[0] + byte(i)
	//	fmt.Printf("%c", zhimu[i])
	//}
	//var intArr [5]int = [...]int{1, -1, 9, 90, 11}
	//var i int = 0
	////for _,x range intArr {
	////    i+=x
	////}
	//for _, v := range intArr {
	//	i += v
	//}
	//fmt.Println(i)
	//fmt.Println(i / len(intArr))
	var intslice [5]int = [...]int{1, 66, 7, 6, 4}
	slice := intslice[1:3] //从第二个元素开始到第四个但不包含第三个所以现在里面只有第二个和第三个
	fmt.Println("intslice ", intslice)
	fmt.Println("元素	", slice)
	fmt.Println("容量", cap(slice))
}
