package cache

import (
	"bitstorm/configs"
	"context"
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	configs.InitConfig()
	res, exists, err := GetRedisCli().Get(context.Background(), "aaaaaaaaaaa")

	// db

	fmt.Println(res, err, err == nil, exists)
}
