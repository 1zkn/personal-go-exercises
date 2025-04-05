package main

import "fmt"

func main(){
	total := 0
	var input int
	for{
		fmt.Println("请输入数字 (0 停止程序)")
		fmt.Scan(&input)
		if(input == 0){
			break
		}
		total = total + input
	}
	fmt.Printf("数字的总和是 %v\n", total)


}