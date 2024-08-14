package main

import "fmt"

func main() {
	// 创建一个通道需要注意,引用类型的数据基本上都是可以使用make来创建的
	ch := make(chan int) //创建了一个无缓冲的,同时通道里面的数据类型是int的通道
	// 向通道里面写入数据?
	// ch <- 10
	go func() {
		// fmt.Println("向通道里面写入数据")
		ch <- 10
	}()
	// 从通道里面读取数据?
	num := <-ch
	fmt.Println("num:", num)
}
