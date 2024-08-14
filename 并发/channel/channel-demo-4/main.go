package main

import "fmt"

func main() {

	ch := make(chan int, 2) //创建一个有缓冲的通道,通道的容量是2
	ch <- 1
	ch <- 2
	// 查看当前通道的长度和容量
	fmt.Println("len(ch):", len(ch))
	fmt.Println("cap(ch):", cap(ch))
	// 输出通道的数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// 查看当前通道的长度和容量
	fmt.Println("len(ch):", len(ch))
	fmt.Println("cap(ch):", cap(ch))

	// <-ch //通道已经空了,再次读取数据会报错

}
