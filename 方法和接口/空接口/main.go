package main

import "fmt"

func makeSound(a interface{}) {
	fmt.Printf("%v\n", a)
}

func main() {

	// 空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。空接口类型的变量可以存储任意类型的变量。

	arr := []int{1, 2, 3}
	makeSound(arr)
}
