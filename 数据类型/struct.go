package main

import "fmt"

type Point struct {
	Xpoint int
	Ypoint int
}

func main() {
	// 使用结构体字面量来初始化结构体变量p
	p1 := Point{1, 2}
	fmt.Printf("X:%d,Y:%d\n", p1.Xpoint, p1.Ypoint)

	// 使用new关键字来初始化
	p2 := new(Point)
	p2.Xpoint = 2
	p2.Ypoint = 3
	fmt.Printf("X:%d,Y:%d\n", p2.Xpoint, p2.Ypoint)

}
