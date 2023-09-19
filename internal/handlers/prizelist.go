package handlers

import (
	service2 "bitstorm/internal/service"
	"context"
	"fmt"
	"sync"
)

var (
	prizeListHandler *PrizeListHandler // 后台管理应用
	once             sync.Once
)

type PrizeListHandler struct {
	ServicePrize service2.PrizeService
}

func initPrizeListHandler() {
	prizeListHandler = &PrizeListHandler{
		ServicePrize: service2.NewPrizeService(),
	}
}

// GetPrizeListHandler 获取奖品列表处理器
func GetPrizeListHandler() *PrizeListHandler {
	once.Do(initPrizeListHandler)
	return prizeListHandler
}

func (a *PrizeListHandler) CheckInput(args ...interface{}) error {
	if len(args) == 0 {
		return nil
	}
	return nil
}

func (a *PrizeListHandler) Process(ctx context.Context, args ...interface{}) (interface{}, error) {
	list, err := a.ServicePrize.GetPrizeList()

	if err != nil {
		return nil, fmt.Errorf("PrizeListHandler|GetPrizeList:%v", err)
	}
	return list, nil
}
