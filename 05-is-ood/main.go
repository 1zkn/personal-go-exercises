package main

import "fmt"

func main() {
	var input int
	fmt.Println("请输入数字")

	fmt.Scan(&input)

	if input%2 == 0 {
		fmt.Println("你输入的是奇数")
	} else {
		fmt.Println("你输入的是偶数")
	}
}
