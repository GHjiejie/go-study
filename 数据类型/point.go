package main

func main() {
	// 声明一个指针
	var p *int
	// 声明一个int类型的变量
	var i int
	// 将i的地址赋值给p
	p = &i
	// 将i的值赋值为1
	i = 1
	// 输出p的值
	println(*p)
	// 将p的值赋值为2
	*p = 2
	// 输出i的值
	println(i)

}
