// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg sync.WaitGroup

// func printHello() {

// 	fmt.Println("协程执行完毕")

// 	defer wg.Done() // 协程执行完毕后，将计数器-1
// }

// // 这里需要解释一下defer函数的作用,这个defer函数会在函数执行完毕后执行，
// // 这里我们在函数执行完毕后将计数器-1，这样我们就可以在主线程中使用sync.WaitGroup的Wait方法等待所有的协程执行完毕后再结束主线程。

// func main() {

// 	wg.Add(1)       // 计数器+1,表明我们主线程需要等待一个协程执行结束后再结束
// 	go printHello() // 创建一个新的协程执行 printHello 函数
// 	// time.Sleep(time.Second) // 等待协程执行完成 这里我们不需要等待了，我们可以使用sync.WaitGroup来等待
// 	wg.Wait() //等待所有的协程执行完毕

// 	fmt.Println("主线程执行完毕")
// }
