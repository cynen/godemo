package gindemo

import (
	"net/http"
	"strings"
)

// Engine 定义一个 Web引擎.
// 参考: https://blog.csdn.net/weixin_41357767/article/details/112581666
type Engine struct {
	// 核心容器,主要是存储handler和req的映射关系.
	// router map[string]http.HandlerFunc
	router map[string]HandlerFunc
	RouterGroup
	groups []*RouterGroup
}

// Default 创建默认Web引擎的方法.
func Default() *Engine {
	// 通过Default ,我们创建一个http 服务器
	engine := New()
	return engine
}

// New 创建引擎.
func New() *Engine {
	engine := &Engine{
		RouterGroup: RouterGroup{
			basePath: "",
		},
	}
	engine.RouterGroup.engine = engine
	engine.groups = append(engine.groups, &engine.RouterGroup)
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
// 此方法是 http.Handler 包中的接口.
// 注意,因为此方法中处理handler关联了 req.Method,所以,我们后续所有handler必须和对应的method绑定.
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 1.原生的http ,未作处理.
	//w.Write([]byte("hello"))

	// 2.标准Web服务,添加handler
	/*
		key := req.Method + "-" + req.RequestURI
		if handler, ok := engine.router[key]; ok {
			handler(w, req)
		} else {
			http.NotFound(w, req)
		}
	*/

	// 3.修改为context
	/*
		context := newContext(w, req)
		context.engine = engine
		key := req.Method + "-" + req.RequestURI
		if handler, ok := engine.router[key]; ok {
			handler(context)
		} else {
			http.NotFound(w, req)
		}
	*/

	// 4.添加中间件.
	context := newContext(w, req)
	context.engine = engine

	// 根据路径判断,这个请求是属于哪个组,然后获取这个组上的中间件.
	// 把这些中间件放到上下文中.
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.basePath) {
			context.handlers = append(context.handlers, group.Handlers...)
		}
	}

	// 把请求处理器也作为"中间件"添加到上下文中.
	key := req.Method + "-" + req.RequestURI
	if handler, ok := engine.router[key]; ok {
		context.handlers = append(context.handlers, handler)
	} else {
		context.handlers = append(context.handlers, func(ctx *Context) {
			http.Error(ctx.Writer, "404 page not found", http.StatusInternalServerError)
		})
	}
	context.Next()
}

// 向engine中添加handlerFunc
// 核心容器,关键就是一个map
func (engine *Engine) addRouter(method string, pattern string, handlerFunc HandlerFunc) {
	if engine.router == nil {
		engine.router = make(map[string]HandlerFunc)
	}
	key := method + "-" + pattern
	engine.router[key] = handlerFunc
}

// Get 这个主要是为了区分GET 和POST的差异. 实际还是调用的 addRouter 方法.
func (engine *Engine) Get(pattern string, handlerFunc HandlerFunc) {
	engine.addRouter("GET", pattern, handlerFunc)
}

// Post 这个主要是为了区分GET 和POST的差异. 实际还是调用的 addRouter 方法.
func (engine *Engine) Post(pattern string, handlerFunc HandlerFunc) {
	engine.addRouter("POST", pattern, handlerFunc)
}
