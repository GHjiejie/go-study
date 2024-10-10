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
