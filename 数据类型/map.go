package main

import "fmt"

func main() {
	// 创建一个map
	map1 := make(map[string]int, 10)
	// 这里创建了一个map，map的key是string类型，value是int类型。map的长度是10，这个长度是map的容量，不是map的长度。
	map1["a"] = 1
	map1["b"] = 2
	map1["c"] = 3
	fmt.Println(map1) // map[a:1 b:2 c:3]

	//判断map里面是否存在某个key
	value, ok := map1["a"]
	if ok {
		fmt.Println(value) // 1
	} else {
		fmt.Println("key not exist")
	}

}
