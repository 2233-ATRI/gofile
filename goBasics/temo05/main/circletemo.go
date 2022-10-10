package main

import "fmt"

//import "fmt"
//
//type clure struct {
//	restw float64
//}
//type methodus struct {
//}
//type masie struct {
//    a int
//	b int
//}
//
//func (a masie) ace()  {
//
//}
//func (c clure) area() float64 {
//	return 3.14 * c.restw * c.restw
//}
//func (b methodus) moersw() {
//	for i := 0; i < 10; i++ {
//		for j := 0; j < 8; j++ {
//			fmt.Print("*")
//		}
//		fmt.Println()
//	}
//}
//func main() {
//	var a clure
//	a.restw = 4.0
//	res := a.area()
//	fmt.Println(res)
//	var b methodus
//	b.moersw()
//}

//type methodutils struct {
//	a [][]int
//}
//
//func (c methodutils) sceew(b int) {
//	for i := 1; i < b; i++ {
//		for j := 1; j < 10; j++ {
//			fmt.Printf("%v * %v = %v", i, j, i*j)
//		}
//		fmt.Scanln()
//	}
//}
//
//func main() {
//	var c methodutils
//	var b int = 6
//	c.sceew(c)
//
//}

type Num struct {
}

func (array Num) Upserver(Aaaay3 [3][3]int) {
	for i := 0; i < len(Aaaay3); i++ {
		for j := 0; j < i; j++ {
			Aaaay3[i][j], Aaaay3[j][i] = Aaaay3[j][i], Aaaay3[i][j]
		}
	}
	fmt.Println(Aaaay3)
}

//
//func (array Num) Upserver2(Aaaay3 [3][3]int) {
//	temparry := [3][3]int{}
//	for i := 0; i < len(Aaaay3); i++ {
//		for j := 0; j < i; j++ {
//			temparry[i][j] = Aaaay3[i][j]
//			Aaaay3[i][j] = Aaaay3[j][i]
//			Aaaay3[j][i] = temparry[i][j]
//		}
//	}
//	fmt.Println(Aaaay3)
//}

func (array Num) vaer(a [3][3]int) {
	var c [3][3]int
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			c[i][j] = a[i][j]
			fmt.Printf("%v", c[i][j])
		}
		fmt.Println()
	}
}

func main() {
	arrinfo := Num{}
	aeey := [3][3]int{
		{0, 1, 2}, /*  第一行索引为 0 */
		{4, 5, 6}, /*  第二行索引为 1 */
		{8, 9, 10}}
	//fmt.Println(aeey)
	//fmt.Println("****")
	arrinfo.Upserver(aeey)
	//arrinfo.Upserver2(aeey)
	arrinfo.vaer(aeey)
}
