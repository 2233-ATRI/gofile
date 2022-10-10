package main

import "fmt"

func main() {
	var day int = 97
	var week int = 7
	var dayweek int = day / week
	var overday int = day % week
	fmt.Println("还有", dayweek, "周,还有", overday, "天")
}
