package main

import "fmt"

func fbn(n int) []uint64 {
	fbnslice := make([]uint64, n)
	fbnslice[0] = 1
	fbnslice[1] = 1
	for i := 2; i < n; i++ {
		if n <= 2 {
			break
		}
		fbnslice[i] = fbnslice[i-1] + fbnslice[i-2]
	}
	return fbnslice
}
func main() {

	slice := fbn(4)
	fmt.Println(slice)
}
