package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {

	for job := range jobs {
		fmt.Println("我处理的数据是：", job)
		time.Sleep(time.Second) // 模拟工作
		results <- job * 2      // 将结果发送到结果通道
	}
}

func main() {
	jobs := make(chan int, 5)    // 创建带缓冲的 jobs 通道
	results := make(chan int, 5) // 创建带缓冲的 results 通道

	// 启动几个 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送一些工作到 jobs 通道
	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs) // 关闭 jobs 通道，表示没有更多的工作

	// 获取结果
	for r := 1; r <= 5; r++ {
		result := <-results
		fmt.Println("Result:", result)
	}
}
