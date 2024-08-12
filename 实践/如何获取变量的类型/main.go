package main

import (
	"fmt"
	"reflect"
)

// 第一种方法就是使用空接口+fmt包里面的%T实现，，当然，你也可以只使用这个fmt包
func OutputVariableType(v interface{}) {
	fmt.Printf("使用fmt包判断变量的类型是：%T\n", v)
}

// 第三中方法是使用类型断言
func OutputType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("这是一个int类型的数据")
	case string:
		fmt.Println("这是一个string类型的数据")
	case [3]int:
		fmt.Println("这是一个元素为int,大小为3的数组")
	case []int:
		fmt.Println("这是一个元素为int的切片")
	}
}

func main() {
	num := 10
	str := "jie"
	arr := [3]int{1, 2, 3}      //注意,这里声明的是一个长度为3的数组,声明数组的话,一定要指定长度
	arr_alice := []int{1, 2, 3} //注意,这里声明的是一个长度不确定的切片,切片和数组不是一个数据结构哈,这个需要我们注意一下,切片是可以动态扩容的,数组长度是不可以改变的

	// 第一种方法就是使用空接口+fmt包里面的%T实现
	OutputVariableType(num)
	OutputVariableType(str)
	OutputVariableType(arr)
	OutputVariableType(arr_alice)
	// 第二种方法就是使用reflect包

	fmt.Println("使用reflect包判断变量的类型是", reflect.TypeOf(num))
	fmt.Println("使用reflect包判断变量的类型是", reflect.TypeOf(str))
	fmt.Println("使用reflect包判断变量的类型是", reflect.TypeOf(arr))
	fmt.Println("使用reflect包判断变量的类型是", reflect.TypeOf(arr_alice))

	// 第三种方法是使用类型开关

	OutputType(num)
	OutputType(str)
	OutputType(arr)
	OutputType(arr_alice)

}
