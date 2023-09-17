package api

import (
	"bitstorm/internal/handlers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// GetPrizeList 获取奖品列表
func GetPrizeList(c *gin.Context) {
	// todo: 参数获取，校验
	rsp := &HttpResponse{}
	handler := handlers.GetFactoryHandler().Get("/get_prize_list")
	prizeList, err := handler.Process(WithReqContext())
	if err != nil {
		log.Errorf("Api|GetPrizeList:%v", err)
		rsp.ResponseWithError(c, CodeGetPrizeInfoErr, err.Error())
	}
	rsp.ResponseWithData(c, prizeList)
}
