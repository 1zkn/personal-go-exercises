package main

import "fmt"

func main() {
	var num int
	fmt.Print("欢迎来到FizzBuzz游戏\n请输入一个数字\n>")
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz ")
		} else if i%3 == 0 {
			fmt.Print("Fizz ")
		} else if i%5 == 0 {
			fmt.Print("buzz ")
		} else {
			fmt.Print(i, " ")
		}
	}
}
