package gindemo

// RouterGroup 实现路由分组
type RouterGroup struct {
	basePath string
	engine   *Engine

	// 存放中间件
	Handlers []HandlerFunc
}

// Use 给分组提供一个Use方法,把中间件存放到对应的RouterGroup中.
func (group *RouterGroup) Use(middlewares ...HandlerFunc) *RouterGroup {
	group.Handlers = append(group.Handlers, middlewares...)
	return group
}

// Group 提供创建 RouterGroup 的函数Group
func (group *RouterGroup) Group(relativePath string) *RouterGroup {
	newGroup := &RouterGroup{
		basePath: relativePath,
		engine:   group.engine,
	}
	group.engine.groups = append(group.engine.groups, newGroup)
	return newGroup
}

// RouterGroup 添加 router,实际也是 添加到容器里.
func (group *RouterGroup) addRouter(method string, pattern string, handlerFunc HandlerFunc) {
	if group.engine.router == nil {
		group.engine.router = make(map[string]HandlerFunc)
	}
	pattern = group.basePath + pattern
	key := method + "-" + pattern
	group.engine.router[key] = handlerFunc
}

// Get 区分GET和POST请求.
func (group *RouterGroup) Get(pattern string, handlerFunc HandlerFunc) {
	group.addRouter("GET", pattern, handlerFunc)
}
func (group *RouterGroup) Post(pattern string, handlerFunc HandlerFunc) {
	group.addRouter("POST", pattern, handlerFunc)
}
