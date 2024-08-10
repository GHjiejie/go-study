package main

import "fmt"

// 声明一个矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// 为这个结构体添加一个方法，由于计算这个矩形的面积
func (r Rectangle) getArea() float64 {
	return r.Height * r.Width
}

func main() {
	// 因为我们为这个结构体定义了一个方法，我们不妨就输出一下这个结构体
	// 初始化一个结构体类型的变量
	rect := Rectangle{4, 5}
	fmt.Printf("rect:%v\n", rect)
	fmt.Printf("rect的矩形的面积是:%f", rect.getArea())
}
