package main

import (
	"godemo/gindemo"
	"net/http"
)

func main() {
	// 1.使用原生的 http 实现 web服务器.
	//gindemo.Main()

	// 2.自定义实现gin框架.
	engine := gindemo.Default()
	engine.Get("/hello", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("hello!"))
	})
	engine.Run()
}
