package main

import "fmt"

func deferExample() {
	x := 10

	// 使用 defer 延迟执行，并将 x 当前的值传递给函数
	defer func(y int) {
		fmt.Println("Value of y:", y)
	}(x) // 这里的 x 是当前的值（10）

	x = 20 // 修改 x 的值
}

func main() {
	deferExample()
}
