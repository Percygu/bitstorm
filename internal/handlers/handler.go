package handlers

type Handler interface {
	CheckInput() error
	Process()
}

// Run 执行函数
func Run(handler Handler) {
	// 1. 参数校验
	err := handler.CheckInput()
	// 校验失败，
	if err != nil {
		return
	}

	// 2. 逻辑处理
	handler.Process()
}
