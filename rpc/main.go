package main

import "rpc/init"

func main() {
	println("hello world")
	// 初始化db
	init.OrmInit();
	// 初始化redis
	// 初始化rpc
	// 初始化日志系统
	// 初始化拦截器
	//
}
