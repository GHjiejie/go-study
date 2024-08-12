package main

import "fmt"

func OutputType(a interface{}) {
	fmt.Printf("该变量的类型是：%T\n", a)
}
func main() {

	number := 10
	OutputType(number)

	str := "hello world"
	OutputType(str)

	arr := [3]int{1, 2, 3}
	OutputType(arr)

}
