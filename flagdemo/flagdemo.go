package flagdemo

import (
	"flag"
	"fmt"
)

// FlagDemo
// 用于读取命令行参数. 生成二进制,进行命令行操作,即可获取参数.
// go build 输出二进制后,再执行 FlagDemo -h
// 参考: https://studygolang.com/articles/21438
func FlagDemo() {

	var username string
	var password string
	var host string
	var port int
	//flag.StringVar(&username, "u", "", "用户名")
	flag.StringVar(&username, "u", "", "用户名")
	flag.StringVar(&password, "p", "", "密码")
	flag.StringVar(&host, "h", "127.0.0.1", "主机")
	flag.IntVar(&port, "P", 3306, "端口")

	flag.Parse()
	fmt.Printf("username=%v password=%v host=%v port=%v\n", username, password, host, port)

}
