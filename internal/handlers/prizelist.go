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
	ServiceAdmin service2.PrizeService
}

func initPrizeListHandler() {
	prizeListHandler = &PrizeListHandler{
		ServiceAdmin: service2.NewPrizeService(),
	}
}

// GetPrizeListHandler 获取全局应用
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
	list, err := a.ServiceAdmin.GetPrizeList()
	if err != nil {
		return nil, fmt.Errorf("AppAdmin|GetPrizeList:%v", err)
	}
	return list, nil
}
