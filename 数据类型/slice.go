package main

import "fmt"

func main() {
	// 切片的声明和数组类似，只是不需要指定长度。切片的长度可以不固定，可以动态增长。切片的声明格式如下：
	// var sliceName []elementType    //需要注意的一个点，就是这个elementType是这个切片中所有元素的类型，而不是切片的类型。
	// 这个和数组声明是一样的，因为我们数组声明 var arrName [arrLength]elementType 也是一样的。

	var slice1 []int
	// 这里声明了一个int类型的切片，切片名为slice1。这个切片的长度是0，因为没有指定长度。这个切片是一个空切片。

	var slice2 = []int{1, 2, 3}
	// 这里声明了一个int类型的切片，切片名为slice2。这个切片的长度是3，因为指定了长度。这个切片是一个非空切片。

	fmt.Println(slice1) // []
	fmt.Println(slice2) // [1 2 3]

	// 你也可以基于一个数组来创建一个切片，这个切片的长度和容量和数组一样。切片的长度是切片中元素的个数，切片的容量是从切片的第一个元素到底层数组的最后一个元素的个数。切片的声明格式如下：
	// sliceName := arrName[start:end]    //这里的arrName是一个数组，start是开始的索引，end是结束的索引。这个切片的长度是end-start，容量是从arrName的start索引到arrName的最后一个元素的个数。
	// 左开右闭区间，

	arr := [5]int{34, 2, 22, 4, 5}

	slice3 := arr[0:3]

	// 针对于某些情况，我们也可以有如下的写法：
	// 比如说，我们想要取出arr的前两个元素，我们可以这样写：
	slice4 := arr[:2]
	// 如果我们要从第二个元素开始取出arr的元素，我们可以这样写：
	alice5 := arr[1:]
	fmt.Println(slice3) // [34 2 22]
	fmt.Println(slice4) // [34 2]
	fmt.Println(alice5) // [2 22 4 5]
}
