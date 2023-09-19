package api

import (
	"bitstorm/internal/handlers"
	"bitstorm/internal/pkg/constant"
	"bitstorm/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PrizeListHandler struct {
	req PrizeListRequest
	// resp     PrizeListResponse
	resp HttpResponse

	// 需要什么Service，就在这里声明
	service service.PrizeService
}

// GetPrizeList 获取奖品列表
func GetPrizeList(c *gin.Context) {
	// todo: 参数获取，校验
	h := PrizeListHandler{
		service: service.NewPrizeService(),
	}
	// HTTP响应
	defer func() {
		// 通过对应的Code，获取Msg
		h.resp.Msg = constant.GetErrMsg(h.resp.Code)
		c.JSON(http.StatusOK, h.resp)
	}()
	// 获取请求数据
	c.ShouldBind(h.req)
	handlers.Run(&h)
}

func (h *PrizeListHandler) CheckInput(ctx context.Context) error {
	h.resp.Code = constant.ErrInputInvalid
	return nil
}

func (h *PrizeListHandler) Process(ctx context.Context) {
	v, err := h.service.GetPrizeList(ctx)
	if err != nil {
		// TODO:
		h.resp.Code = constant.PrizeStatusDelete
		// log.Errorf()
		return
	}

	// 继续处理
	h.resp.Data = v
	return
}
