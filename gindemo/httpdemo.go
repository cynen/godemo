package gindemo

import "net/http"

// 使用原生的 net 构建http服务.

func Main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8888", nil)
}

func hello(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("hello"))
}
