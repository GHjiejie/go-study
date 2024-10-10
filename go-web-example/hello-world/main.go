package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 下面介绍返回响应的几种方式：
	// 1. 使用 fmt.Fprintf 格式化字符串
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)

	// 2. 使用 io.WriteString 写入字符串
	w.Write([]byte("大家好，我是Jie!!!"))

	// 3. 使用 json.NewEncoder 返回 JSON 格式数据
	response := map[string]interface{}{"code": 200, "msg": "success", "data": "Hello, you've requested: " + r.URL.Path}
	json.NewEncoder(w).Encode(response)

}
func main() {
	// 注册请求处理函数
	http.HandleFunc("/hello", helloHandler)
	// 监听端口，同时输出消息
	http.ListenAndServe(":8080", nil)
}
