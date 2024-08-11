package main

import "fmt"

// 声明一个接口
// 首先我们需要知道接口的定义是什么：在go里面，接口是一直抽象类型，他定义了一组方法签名。
// 他不关心方法的具体实现的细节，只对方法名，方法接受的参数，以及方法的返回值类型做了限制
type Shape interface {
	Area() float64      //图形的面积
	Perimeter() float64 //图形的周长
}

// 声明一个矩形结构体
type Rect struct {
	Width  float64
	Height float64
}

// 下面就是对接口的实现，为结构体里面添加接口对应的方法
func (r Rect) Area() float64 {
	return r.Height * r.Width
}

// 矩形的周长
func (r Rect) Perimeter() float64 {
	return (r.Height + r.Width) * 2
}

// 在声明一个圆形的结构体
type Circle struct {
	Radius float64
}

// 圆形的面积
func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// 圆形的周长
func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// 输出图形的面积与周长

func Output(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimeter: %f\n", s.Perimeter())
}

func main() {

	r := Rect{Width: 5, Height: 10}
	c := Circle{Radius: 5}

	// 因为此时我们的这个Rect已经实现了这个接口，所以我们就可以执行这个Output方法了
	Output(r)

	// 同理，我们的Circle也实现了这个接口，所以我们也可以执行这个Output方法
	Output(&c)
}
