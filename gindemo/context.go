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

	// 添加中间件.
	handlers []HandlerFunc
	// 中间件的索引
	index int
}

// HandlerFunc 这个是为了使用context进行处理.
// 替换原生的 http.HandlerFunc
type HandlerFunc func(ctx *Context)

// newContext 新建一个context
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: req,
		index:   -1,
	}
}

// Next 这个是整个责任链的核心方法.每当这个方法被调用,就执行下一个中间件函数.
func (c *Context) Next() {
	c.index++
	for c.index < len(c.handlers) {
		// 执行下一个中间件函数
		c.handlers[c.index](c)
		c.index++
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
