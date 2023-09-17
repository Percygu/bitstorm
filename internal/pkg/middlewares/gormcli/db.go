package gormcli

import (
	"bitstorm/configs"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
	"sync"
	"time"
)

type ctxTransactionKey struct {
}

var (
	db     *gorm.DB
	dbOnce sync.Once
)

// openDB 连接db
func openDB() {
	dbConfig := configs.GetGlobalConfig().DbConfig
	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.User,
		dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DataBase)
	// log.Info("mdb addr:" + connArgs)

	var err error
	db, err = gorm.Open(mysql.Open(connArgs), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("fetch gormcli connection err:" + err.Error())
	}

	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConn)                                        // 设置最大空闲连接
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConn)                                        // 设置最大打开的连接
	sqlDB.SetConnMaxLifetime(time.Duration(dbConfig.MaxIdleTime * int64(time.Second))) // 设置空闲时间为(s)
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	dbOnce.Do(openDB)
	return db
}

func CloseDB() {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Errorf("close gormcli err:%v", err)
		}
		sqlDB.Close()
	}
}

func Transaction(ctx context.Context, fc func(txctx context.Context) error) error {
	db := GetDB().WithContext(ctx)

	return db.Transaction(func(tx *gorm.DB) error {
		txctx := CtxWithTransaction(ctx, tx)
		return fc(txctx)
	})
}

func GetDBFromCtx(ctx context.Context) *gorm.DB {
	iface := ctx.Value(ctxTransactionKey{})

	if iface != nil {
		tx, ok := iface.(*gorm.DB)
		if !ok {
			log.Panicf("unexpect context value type: %s", reflect.TypeOf(tx))
			return nil
		}

		return tx
	}

	return GetDB().WithContext(ctx)
}

func CtxWithTransaction(ctx context.Context, tx *gorm.DB) context.Context {
	c := context.WithValue(ctx, ctxTransactionKey{}, tx)
	return c
}
