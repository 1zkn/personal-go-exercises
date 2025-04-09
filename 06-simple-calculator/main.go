package main

import "fmt"

func main() {
	var num1, num2, op float64
	fmt.Println("简单计算器")
	fmt.Print("输入的一个数字:")
	fmt.Scan(&num1)
	fmt.Print("输入的二个数字:")
	fmt.Scan(&num2)
	fmt.Print("请选择您的运算符\n1 - 加法\n2 - 减法\n3 - 乘法\n4 - 除法\n>")
	fmt.Scan(&op)
	switch op {
	case 1:
		fmt.Printf("%f + %f = %f", num1, num2, num1+num2)
	case 2:
		fmt.Printf("%f - %f = %f", num1, num2, num1-num2)
	case 3:
		fmt.Printf("%f * %f = %f", num1, num2, num1*num2)
	case 4:
		if num2 != 0 {
			fmt.Printf("%f / %f = %f", num1, num2, num1/num2)
		} else {
			fmt.Println("第二个数字不能是0哦")
		}
	default:
		fmt.Println("请输入可用的选项")
	}
}
