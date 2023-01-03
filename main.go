package main

import (
	"godemo/sqldemo"
	"godemo/syncdemo"
)

func main() {
	// 1.测试GinWeb
	GinWebDemo()

	// 2.测试sql
	//SqlDemo()

	// 3.SyncDemo 测试sync包
	// Syncdemo()
	syncdemo.TestCond()
	// 4.获取命令行参数.
	//flagdemo.FlagDemo()

	// 5.Chaneldemo

}

func GinWebDemo() {
	// 1.使用原生的 http 实现 web服务器.
	//gindemo.Main()

	// 2.实现最基本的Web
	//gindemo.TestWeb()

	// 3.添加分组功能.
	//gindemo.TestGroup()

	// 4.测试Context
	//gindemo.TestContext()
}

func SqlDemo() {
	sqldemo.TestSql1()
	sqldemo.TestSql2()
}

func Syncdemo() {
	syncdemo.TestMap()
	syncdemo.SyncMap()
	syncdemo.WgTest1()
}
