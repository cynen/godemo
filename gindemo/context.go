package gindemo

import (
	"encoding/json"
	"net/http"
)

// H 用来封装整个用于转换json和map,减少代码量.
// Gin框架就是这么干的.
type H map[string]interface{}

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	engine  *Engine
}

// HandlerFunc 这个是为了使用context进行处理.
// 替换原生的 http.HandlerFunc
type HandlerFunc func(ctx *Context)

// newContext 新建一个context
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: req,
	}
}

// JSON 封装响应 JSON的方法.
func (context *Context) JSON(code int, obj interface{}) {
	context.Writer.WriteHeader(code)
	context.Writer.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(context.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(context.Writer, err.Error(), http.StatusInternalServerError)
	}
}
