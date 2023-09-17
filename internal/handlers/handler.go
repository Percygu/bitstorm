package handlers

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type Handler interface {
	CheckInput(args ...interface{}) error
	Process(ctx context.Context, args ...interface{}) (interface{}, error)
}

type Factory struct {
	handlerMap map[string]Handler
}

var factory *Factory

func (f *Factory) Register(key string, handler Handler) error {
	if _, ok := f.handlerMap[key]; ok {
		return fmt.Errorf("handler registered,key is %s", key)
	}
	f.handlerMap[key] = handler
	return nil
}

func (f *Factory) Unregister(key string) error {
	if _, ok := f.handlerMap[key]; !ok {
		return fmt.Errorf("handler not registered,key is %s", key)
	}
	delete(f.handlerMap, key)
	return nil
}

func (f *Factory) Get(key string) Handler {
	var (
		ok      bool
		handler Handler
	)
	if handler, ok = f.handlerMap[key]; !ok {
		log.Errorf("handler registered,key is %s", key)
		return nil
	}
	return handler
}

func GetFactoryHandler() *Factory {
	return factory
}

func InitHandlerFactory() {
	factory = &Factory{
		handlerMap: make(map[string]Handler),
	}
	factory.Register("/get_prize_list", GetPrizeListHandler())
}
