package utils

import "fmt"

func Cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
		break
	case '-':
		res = n1 - n2
		break
	case '*':
		res = n1 * n2
		break
	case '/':
		res = n1 / n2
		break
	default:
		fmt.Println("error")
	}
	return res
}

//为了让其他的包也可以使用Cal这个函数，需要将其中的C进行大写，起作用相当于其他语言当中的public
