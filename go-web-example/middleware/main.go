package main

import (
	"fmt"
	"log"
	"net/http"
)

// 因为是中间件函数,所以我们最好将接受的参数命名为next或者nextHandler,这样可以清晰的表明这是一个中间件函数(按照业内的规范来)
func logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("请求URL:", r.URL.Path)
		next(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "输出此时匹配中间件的的路径+%s/n", r.URL.Path)
	w.Write([]byte("我是foo"))
}

func bar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("我是bar"))
}

func main() {
	// 我们使用默认的路由,中间件函数
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))
	http.ListenAndServe(":8080", nil)
}
