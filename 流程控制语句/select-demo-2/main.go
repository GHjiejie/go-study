package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// 模拟从通道接收数据
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 42
	}()

	for i := 0; i < 5; i++ {
		select {
		case val := <-ch:
			fmt.Println("Received:", val)
		default:
			fmt.Println("No value received, doing something else...")
			time.Sleep(500 * time.Millisecond) // 做一些其他事情
		}
	}
}
