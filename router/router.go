package router

import (
	"bitstorm/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

const SessionKey = "lottery_session" // 鉴权session

// AuthMiddleWare 鉴权中间件
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if session, err := c.Cookie(SessionKey); err == nil {
			if session != "" {
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		c.Abort()
		return
	}
}

func setRoutes(r *gin.Engine) {
	setAdminRoutes(r)
	setLotteryRoutes(r)
}

func setAdminRoutes(r *gin.Engine) {
	adminGroup := r.Group("admin")
	// 获取奖品列表
	adminGroup.GET("/get_prize_list", api.GetPrizeList)
}

func setLotteryRoutes(r *gin.Engine) {
	lotteryGroup := r.Group("lottery")
	// 获取中奖
	lotteryGroup.GET("/get_lucky", AuthMiddleWare(), api.GetPrizeList)
}
