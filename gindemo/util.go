package gindemo

import (
	"log"
	"time"
)

func resolveAddress(addr []string) string {
	switch len(addr) {
	case 0:
		return ":8080"
	case 1:
		return addr[0]
	default:
		panic("Too Many Parameters")
	}
}

func TestWeb() {
	// 2.自定义实现gin框架.
	// 参考: https://blog.csdn.net/weixin_41357767/article/details/112581666
	/*
		engine := Default()
		engine.Get("/hello", func(w http.ResponseWriter, request *http.Request) {
			w.Write([]byte("hello!"))
		})
		engine.Run(":8888")
	*/

}
func TestGroup() {
	// 3. 实现分组的gin
	// 新建引擎后,我们需要初始化我们的group 路径.
	// 这里是因为 engine内部有一个匿名的结构体: RouterGroup
	// 内嵌结构体特殊一点，我们可以通过父结构体直接拿到内嵌结构体的字段，而不需要通过内嵌结构体间接获得
	// 参考: https://www.bilibili.com/read/cv13062515

	/*
		engine := Default()

		userGroup := engine.Group("/user")

		userGroup.Get("/login", func(w http.ResponseWriter, request *http.Request) {
			w.Write([]byte("login Success"))
		})
		userGroup.Post("/logout", func(w http.ResponseWriter, request *http.Request) {
			w.Write([]byte("logout sucess"))
		})
		engine.Run(":8888")
	*/
}

func TestContext() {

	// 4. 因为使用了context,所有的handlerFunc都需要变更.所以上面的demo需要修改.
	// 把一系列注册路由方法的参数改为 Context，为此我们还要重新封装一下处理方法，不能再用 http.HandlerFunc了
	engine := Default()
	//engine.Get("/ping", func(ctx *Context) {
	//	ctx.JSON(200, H{
	//		"message": "pong",
	//	})
	//})

	userGroup := engine.Group("/user")

	userGroup.Get("/login", func(ctx *Context) {
		ctx.JSON(200, H{
			"massage": "login OK",
		})
	})

	userGroup.Post("/logout", func(ctx *Context) {
		ctx.JSON(200, H{
			"message": "logout",
		})
	})

	engine.Run(":8888")
}

func TestMiddler() {
	engine := Default()
	engine.Use(Timer)
	engine.Get("/ping", func(ctx *Context) {
		ctx.JSON(200, H{
			"message": "OK",
		})
	})
	engine.Run(":8888")
}

// Timer 自定义一个中间件,用于记录请求响应时间.
func Timer(ctx *Context) {
	t := time.Now()
	ctx.Next()
	log.Printf("use time: %v\n", time.Since(t))
	//fmt.Println("use time : ", time.Since(t))
}
