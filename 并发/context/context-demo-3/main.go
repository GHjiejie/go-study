package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// var wg sync.WaitGroup

// // processRequest 模拟处理请求
// func processRequest(userID int, requestType string, status string) {
// 	defer wg.Done()             //一个job工作完成后就调用Done方法，表示这个job已经完成了
// 	time.Sleep(2 * time.Second) // 模拟处理时间
// 	fmt.Printf("Processing request for UserID: %d, Type: %s, Status: %s\n", userID, requestType, status)
// }

// func main() {
// 	// 模拟多个用户请求
// 	for i := 1; i <= 3; i++ {
// 		wg.Add(1)                             // 每个请求增加一个 WaitGroup
// 		go processRequest(i, "GET", "Active") // 显式传递多个参数
// 	}

// 	wg.Wait()
// }

// 我们使用context来改进上面的代码，代码如下：

func processRequest(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	// 从上下文中获取值
	userID := ctx.Value("userID").(int) // 使用强制类型转换
	requestType := ctx.Value("requestType").(string)
	status := ctx.Value("status").(string)

	time.Sleep(2 * time.Second) // 模拟处理时间
	fmt.Printf("Processing request for UserID: %d, Type: %s, Status: %s\n", userID, requestType, status)
}

func main() {
	var wg sync.WaitGroup

	// 创建根上下文
	for i := 1; i <= 3; i++ {
		// ctx := context.Background()
		// 初始化ctx
		ctx := context.Background()
		// 在上下文中存储请求信息
		ctx = context.WithValue(ctx, "userID", i)
		ctx = context.WithValue(ctx, "requestType", "GET")
		ctx = context.WithValue(ctx, "status", "Active")

		wg.Add(1)
		go processRequest(ctx, &wg) // 只需传递上下文
	}

	wg.Wait()
}
