package main

import (
	"fmt"
	"sync"
	"time"
)

// download 模拟下载一个文件
func download(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("开始下载 %s...\n", url)
	time.Sleep(time.Second) // 模拟下载过程
	fmt.Printf("%s 下载完成！\n", url)
}

func main() {
	urls := []string{
		"https://example.com/file1.txt",
		"https://example.com/file2.jpg",
		"https://example.com/file3.pdf",
		"https://example.com/file4.zip",
		"https://example.com/file5.mp4",
	}

	var wg sync.WaitGroup
	wg.Add(len(urls)) // 设置需要等待的协程数量

	// 使用协程并发下载文件
	for _, url := range urls {
		go download(url, &wg)
	}

	wg.Wait() // 等待所有下载任务完成
	fmt.Println("所有文件下载完成！")
}
