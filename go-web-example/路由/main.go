package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("你好,接下来是使用mux路由模块处理请求: " + r.URL.Path))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler)
	// 如果我们不使用mux模块的话,那么下面代码里面就会是这样 http.ListenAndServe(":8080", nil)
	// nil代表使用的是默认路由器(也就是net/http),如果要使用自己的路由器,那么就需要将nil替换成为自己的路由器变量
	http.ListenAndServe(":8080", r)
}
