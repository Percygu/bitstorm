package handlers

import (
	"bitstorm/internal/pkg/constant"
	"bitstorm/internal/pkg/utils"
	"context"
)

type Handler interface {
	CheckInput(ctx context.Context) error
	Process(ctx context.Context)
}

// Run 执行函数
func Run(handler Handler) {
	ctx := context.WithValue(context.Background(), constant.ReqID, utils.NewUuid())
	// 1. 参数校验
	err := handler.CheckInput(ctx)
	// 校验失败，
	if err != nil {
		return
	}

	// 2. 逻辑处理
	handler.Process(ctx)
}
