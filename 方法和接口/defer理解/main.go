package main

import "fmt"

func riskyOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("About to cause a panic...")
	panic("something went wrong") // 触发 panic
}

func main() {
	riskyOperation()
	fmt.Println("Program continues...")
}
