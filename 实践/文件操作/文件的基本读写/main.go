package main

import (
	"fmt"
	"os"
)

func main() {

	// 首先创建一个文件，然后写入内容，最后读取文件内容。
	// 创建文件
	file, err := os.Create("test.txt")
	// 判断错误(err)
	if err != nil {
		fmt.Println("创建文件发生了错误", err)
		return
	}
	// 延迟关闭文件
	defer file.Close()

	// 接下来就是向文件写入内容

	// file.WriteString("Hello, world!")          // 写入内容,但是我们并没有检查错误,所以这里可能会出错,我们最好使用逗号运算符来检查错误
	_, err = file.WriteString("大家好,我是jie") // 写入内容,并检查错误,
	if err != nil {
		fmt.Println("写入过程发生了错误", err)
		return
	}

	// 最后就是读取文件内容
	// 读取文件的全部内容
	content, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("读取文件内容发生了错误", err)
		return
	}

	fmt.Println("文件内容是(ASCII码):", content) //这个输出的是ASCII码,我们可以使用string()函数将其转换为字符串
	fmt.Println("文件内容是(字符串形式):", string(content))

	// 向文件追加内容
	file, err = os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("打开文件发生了错误", err)
		return
	}

	// 执行向文件追加内容的操作
	_, err = file.WriteString("我是追加的内容")
	if err != nil {
		fmt.Println("追加内容发生了错误", err)
		return
	}

}
