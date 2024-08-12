// package main

// import (
// 	"fmt"
// 	"time"
// )

// func printHello() {
// 	time.Sleep(time.Second * 2) //等待2s,但是我们的主线程只等待1s，主死从随，主线程1s后就结束了，子线程就不会再执行了
// 	// 那么我们不知道子线程到底什么时候执行结束，我们应该怎么去处理呢？也就是说,我们的主线程应该等待子线程执行完毕后再结束
// 	fmt.Println("Hello, goroutine!")
// }

// func main() {
// 	go printHello()         // 创建一个新的协程执行 printHello 函数
// 	time.Sleep(time.Second) // 等待协程执行完成
// }
