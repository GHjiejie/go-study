package main

import (
	"fmt"
	"time"
)

// 生产者
func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Printf("Produced: %d\n", i)
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

// 消费者
func consumer(ch <-chan int) {
	for i := range ch {
		fmt.Printf("Consumed: %d\n", i)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int, 2) // 设置缓冲区大小为 2

	go producer(ch)
	consumer(ch)
}
