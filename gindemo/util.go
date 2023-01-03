package gindemo

import "net/http"

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
	engine := Default()
	engine.Get("/hello", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("hello!"))
	})
	engine.Run(":8888")
}
func TestGroup() {
	// 3. 实现分组的gin
	engine := Default()
	// 新建引擎后,我们需要初始化我们的group 路径.
	// 这里是因为 engine内部有一个匿名的结构体: RouterGroup
	// 内嵌结构体特殊一点，我们可以通过父结构体直接拿到内嵌结构体的字段，而不需要通过内嵌结构体间接获得
	// 参考: https://www.bilibili.com/read/cv13062515
	userGroup := engine.Group("/user")

	userGroup.Get("/login", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("login Success"))
	})
	userGroup.Post("/logout", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("logout sucess"))
	})
	engine.Run(":8888")
}
