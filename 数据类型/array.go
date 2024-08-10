package main

import "fmt"

func main() {

	// 使用数组字面量
	var arr1 = [3]int{1, 2, 3} // 先声明数组
	fmt.Println(arr1)          // [1 2 3]

	// 使用:=声明数组
	arr2 := [3]string{"hello", "world", "!"}
	fmt.Println(arr2) // [hello world !]

}
