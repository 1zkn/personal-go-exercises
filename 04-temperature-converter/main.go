package main

import "fmt"

func main(){
	var input float64
	fmt.Print("输入摄氏度温度: ")

	fmt.Scan(&input)

	fmt.Println("转换为华氏度是", input * 9/5 + 32)
}