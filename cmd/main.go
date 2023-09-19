package main

import (
	"bitstorm/configs"
	"bitstorm/router"
)

func Init() {
	configs.InitConfig()
	// gin handler 初始化
}

func main() {
	Init()
	router.InitRouterAndServe()
}
