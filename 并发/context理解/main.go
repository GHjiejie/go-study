package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("任务完成")
	case <-ctx.Done():
		fmt.Println("任务被取消:", ctx.Err())
	}
}

func main() {
	// 创建一个带有超时时间的 context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// 启动一个 goroutine 执行长时间任务
	go longRunningTask(ctx)

	// 等待任务完成或超时
	time.Sleep(3 * time.Second)
}
