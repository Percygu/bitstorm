package router

import (
	"bitstorm/configs"
	"bitstorm/internal/pkg/middlewares/log"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
)

// InitRouterAndServe 路由配置、启动服务
func InitRouterAndServe() {
	setAppRunMode()
	r := gin.Default()

	setMiddleWare(r)
	// 设置路由
	setRoutes(r)

	// 启动server
	port := configs.GetGlobalConfig().AppConfig.Port
	if err := r.Run(":" + strconv.Itoa(port)); err != nil {
		log.Errorf("start server err:" + err.Error())
	}
}

// setAppRunMode 设置运行模式
func setAppRunMode() {
	if configs.GetGlobalConfig().AppConfig.RunMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func setGinLog(out io.Writer) {
	gin.DefaultWriter = out
	gin.DefaultErrorWriter = out
}

// 设置一些中间件：比如Recover

func setMiddleWare(r *gin.Engine) {
	r.Use()
}
