package main

import (
	"godemo/gindemo"
	"net/http"
)

func main() {
	//gindemo.Main()
	engine := gindemo.Default()
	engine.Get("/hello", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("hello!"))
	})
	engine.Run()
}
