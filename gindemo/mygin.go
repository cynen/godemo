package gindemo

import (
	"net/http"
)

// Engine 定义一个 Web引擎.
type Engine struct {
	// 核心容器,主要是存储handler和req的映射关系.
	router map[string]http.HandlerFunc
}

// Default 创建默认Web引擎的方法.
func Default() *Engine {
	// 通过Default ,我们创建一个http 服务器
	engine := New()
	return engine
}

// New 创建引擎.
func New() *Engine {
	engine := &Engine{}
	return engine
}

// Run 启动Web服务器
func (engine *Engine) Run(addr ...string) error {
	address := resolveAddress(addr)
	// engine 要能够被此方法接收为参数,必须实现 此接口的方法. http.Handler
	return http.ListenAndServe(address, engine)
}

// ServeHTTP 实现了ServeHTTP方法就是实现了 Handler 接口.
// 只有实现了此方法的handler,才可以被Run方法中的 http.ListenAndServe 接收为参数.
// 此方法是 http.Handler 包中的 接口.
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//w.Write([]byte("hello"))
	key := req.Method + "-" + req.RequestURI
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		http.NotFound(w, req)
	}
}

// 向engine中添加handlerFunc
// 核心容器,关键就是一个map
func (engine *Engine) addRouter(method string, pattern string, handlerFunc http.HandlerFunc) {
	if engine.router == nil {
		engine.router = make(map[string]http.HandlerFunc)
	}
	key := method + "-" + pattern
	engine.router[key] = handlerFunc
}

// Get 这个主要是为了区分GET 和POST的差异. 实际还是调用的 addRouter 方法.
func (engine *Engine) Get(pattern string, handlerFunc http.HandlerFunc) {
	engine.addRouter("GET", pattern, handlerFunc)
}

// Post 这个主要是为了区分GET 和POST的差异. 实际还是调用的 addRouter 方法.
func (engine *Engine) Post(pattern string, handlerFunc http.HandlerFunc) {
	engine.addRouter("POST", pattern, handlerFunc)
}
