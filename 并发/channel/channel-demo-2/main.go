package main

import "fmt"

func main() {
	ch := make(chan int) //创建一个无缓冲同时接受类型为int的channel

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Print("send: ", i, "\n")
			ch <- i
		}
		defer close(ch)
	}()

	// 使用range遍历channel
	for v := range ch {
		fmt.Println("receive: ", v)
	}
}
