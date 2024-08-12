package main

import (
	"fmt"
	"time"
)

func printHello(done chan bool) {

	time.Sleep(time.Second * 2)
	fmt.Println("Hello, goroutine!")
	done <- true // 发送信号，表示协程已完成
}

func main() {
	done := make(chan bool)
	go printHello(done)
	<-done // 阻塞主线程，直到接收到信号
	fmt.Println("main 函数结束")

}
