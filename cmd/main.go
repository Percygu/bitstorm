package main

import (
	"bitstorm/configs"
	"bitstorm/internal/handlers"
	"bitstorm/router"
)

func Init() {
	configs.InitConfig()
	handlers.InitHandlerFactory()
}

func main() {
	Init()
	router.InitRouterAndServe()
}
