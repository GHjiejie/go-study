package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义命令行标志
	name := flag.String("name", "World", "a name to say hello to")
	age := flag.Int("age", 0, "your age")
	verbose := flag.Bool("v", false, "enable verbose mode")

	// 解析命令行参数
	flag.Parse()

	// 打印解析后的标志值
	if *verbose {
		fmt.Println("Verbose mode enabled")
	}
	fmt.Printf("Hello, %s! You are %d years old.\n", *name, *age)
}
