package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建文件夹
	// 首先判断文件夹是否存在，如果不存在则创建文件夹
	// 创建文件夹的方法是 os.MkdirAll()，该方法会递归创建文件夹，如果文件夹已经存在，则不会报错
	// 判断文件夹是否存在
	_, err := os.Stat("test")
	if err != nil {
		// 如果文件夹不存在，则创建文件夹
		err := os.MkdirAll("test", os.ModePerm)
		if err != nil {
			fmt.Println("创建文件夹发生了错误", err)
			return
		} else {
			fmt.Println("文件夹创建成功")
		}
	} else {
		fmt.Println("文件夹已经存在")
	}

}
