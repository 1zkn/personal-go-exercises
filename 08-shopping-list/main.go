package main

import "fmt"

type Articles struct {
	name     string
	priority string
}

func main() {
	var shopList []Articles

	for {
		fmt.Println("\n====== 购物清单操作菜单 ======")
		fmt.Println("1. 添加商品")
		fmt.Println("2. 删除商品")
		fmt.Println("3. 显示清单")
		fmt.Println("0. 退出")
		fmt.Print("> ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addItem(&shopList)
		case 2:
			removeItem(&shopList)
		case 3:
			printArticles(shopList)
		case 0:
			fmt.Println("再见喵～祝你购物愉快！🐱🛒")
			return
		default:
			fmt.Println("无效选项，请重新输入喵～")
		}
	}
}

func printArticles(a []Articles) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("[id: %v] 商品: %v, 优先级: %v\n", i, a[i].name, a[i].priority)
	}

}

func addItem(shopList *[]Articles) {
	var inputName, inputPriority string
	fmt.Print("请输入商品的名称: ")
	fmt.Scan(&inputName)
	fmt.Print("表明优先级: ")
	fmt.Scan(&inputPriority)
	*shopList = append(*shopList, Articles{name: inputName, priority: inputPriority})
}

func removeItem(shopList *[]Articles) {
	printArticles(*shopList)
	var inputId int
	fmt.Print("输入你想删除的商品id: ")
	fmt.Scan(&inputId)
	if inputId >= len(*shopList) || inputId < 0 {
		fmt.Println("无效的ID!")
		return
	}
	*shopList = append((*shopList)[:inputId], (*shopList)[:inputId+1]...)
}
