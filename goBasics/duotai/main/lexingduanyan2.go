package main

import "fmt"

func typeJunpe(items ...interface{}) {
	for idx, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v", idx, x)
			fmt.Println()
		case int:
			fmt.Printf("第%v个参数是Int类型，值是%v", idx, x)
			fmt.Println()
		case float64:
			fmt.Printf("第%v个参数是float64类型，值是%v", idx, x)
			fmt.Println()
		case float32:
			fmt.Printf("第%v个参数是float32类型，值是%v", idx, x)
			fmt.Println()
		case string:
			fmt.Printf("第%v个参数是String类型，值是%v", idx, x)
			fmt.Println()
		default:
			fmt.Printf("第%v个参数不确定，值是%v", idx, x)
		}
	}
}
func main() {
	var n1 float64 = 1.9
	var n2 float64 = 3.98
	var n3 int = 19
	var Name = "tom"
	typeJunpe(n1, n2, n3, Name)
}
