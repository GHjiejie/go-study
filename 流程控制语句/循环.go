package main

import "fmt"

func main() {
	// 使用for循环和索引遍历
	arr := [...]int{13, 4, 6, 3, 1, 2, 5, 6}
	//我们使用...这个操作符是因为这个数组的长度会根据后面的元素数量动态变换，当然你也可可以指定长度2
	for index := 0; index < len(arr); index++ {
		fmt.Println(arr[index])
	}

	// 使用while循环，虽然go语言里面没有这个关键字，但是我们的for就可以实现这个功能

	// var i int = 0
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// 也可以使用for  range循环
	// 语法是： for index, value := range collection {
	// 	 ...
	// }
	// index数数组元素的下标，value是该数组元素下标对应的值
	// 当然，我们也可以仅处理下标或者值，我们只需要将他们设置为_即可
	// collection是集合的值，不仅限于数组，还有map数据结构这些
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// 忽略集合的下标，因为我们只需要值
	for _, value := range arr {
		fmt.Printf("value:%d\n", value)
	}

	for index, _ := range arr {
		fmt.Printf("index: %d\n", index)
	}

}
