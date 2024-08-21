package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-shutdown // 等待信号
		fmt.Println("Received shutdown signal.")
		// 这里可以执行清理工作
		time.Sleep(2 * time.Second) // 模拟清理过程
		fmt.Println("Cleanup done. Exiting.")
		os.Exit(0)
	}()

	// 主程序逻辑
	fmt.Println("Running... (press Ctrl+C to exit)")
	for {
		time.Sleep(1 * time.Second) // 模拟工作
	}
}
