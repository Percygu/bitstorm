package router

import (
	"bitstorm/configs"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"strconv"
)

// InitRouterAndServe 路由配置、启动服务
func InitRouterAndServe() {
	setAppRunMode()
	r := gin.Default()

	// 设置路由
	setRoutes(r)

	// 启动server
	port := configs.GetGlobalConfig().AppConfig.Port
	if err := r.Run(":" + strconv.Itoa(port)); err != nil {
		log.Error("start server err:" + err.Error())
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
