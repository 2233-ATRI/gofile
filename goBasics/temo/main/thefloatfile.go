package main

import "fmt"

func main() {
	var num1 float64 = -0.188634
	var num2 float32 = -1.291553
	fmt.Println("num1", num1, "num2", num2)
	var num3 float32 = 193.000014
	var num4 float64 = 193.000014
	fmt.Println(num3, "\n", num4)
	var num5 = 1.23
	fmt.Printf("nums %T", num5)
	var num6 float64 = 5.1255e2
	fmt.Printf("num6=", num6)
	//float中64比32位精度更高
}
