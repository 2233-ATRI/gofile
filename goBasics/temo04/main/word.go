package main

import "fmt"

func main() {
	//var intslice [5]int = [...]int{1, 66, 7, 6, 4}
	//slice := intslice[1:3] //从第二个元素开始到第四个但不包含第三个所以现在里面只有第二个和第三个//第一个种方法
	//fmt.Println("intslice ", intslice)
	//fmt.Println("元素	", slice)
	//fmt.Println("容量", cap(slice))
	//var clice []float64 = make([]float64, 10, 10)//第一个10是大小，第二个10是它的容量，大小不可以大于容量
	//fmt.Println(clice)//第二种办法
	//var slice []string = []string{"tom", "jeck"}//第三个方法
	//fmt.Println(len(slice))
	//fmt.Println(cap(slice))
	//var arr []int = []int{10, 20, 30, 40, 50}
	//slice := arr[1:4]
	//for i := 0; i < len(slice); i++ {
	//	fmt.Printf("slice[%v] = %v ", i, slice[i])
	//
	//}
	//
	//var slice3 []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	//slice3 = append(slice3, 20, 30, 50)
	//slice3 =append(slice3,slice3...)//追加切片的固定写法
	//fmt.Println(slice3)
	//切片append是对原先的数组进行扩容
	//go底层会创造一个新的数组newArr，原来的元素会重新引用到新的数组当中

	//newArr是不可见的
	var slice4 []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	var slice5 []int = make([]int, 10)
	copy(slice4, slice5) //把5中的copy给4
	fmt.Println(slice4)

}
