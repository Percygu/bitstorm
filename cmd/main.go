package main

import (
	"bitstorm/configs"
	"bitstorm/router"
)

func Init() {
	configs.InitConfig()
}

func main() {
	Init()
	router.InitRouterAndServe()
}
