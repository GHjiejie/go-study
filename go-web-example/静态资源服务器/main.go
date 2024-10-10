package main

import (
	"fmt"
	"net/http"
)

func handleStatic(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "接下来是处理静态资源的代码;%s\n", r.URL.Path)
}
func main() {
	http.HandleFunc("/", handleStatic)
	// 设置静态资源目录
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}

// 以static目录开始的请求将会被重定向到static目录下的文件，例如请求/static/test.jpg将会被重定向到static目录下的test.jpg文件。
// 示例请求:http://localhost:8080/static/test.jpg
// 示例请求:http://localhost:8080/static/test.txt
