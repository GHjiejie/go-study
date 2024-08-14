package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	fmt.Println("输出ctx:", ctx)
	select {
	case <-time.After(2 * time.Second): // 模拟一些工作
		fmt.Println("Work completed")
	case <-ctx.Done(): // 响应取消信号
		fmt.Println("Work canceled:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second) // 设置超时
	defer cancel()                                                          // 确保在 main 函数退出时调用 cancel

	go doWork(ctx) // 启动同伴 goroutine

	time.Sleep(3 * time.Second) // 等待一段时间以演示
}
