package main

import (
	"bitstorm/configs"
	"bitstorm/internal/pkg/middlewares/cache"
	"bitstorm/internal/pkg/middlewares/gormcli"
	"bitstorm/internal/pkg/middlewares/log"
	"bitstorm/router"
)

func Init() {
	conf := configs.InitConfig()
	logConf := conf.LogConfig
	dbConf := conf.DbConfig
	cacheConf := conf.RedisConfig

	// 初始化日志
	log.Init(
		log.WithFileName(logConf.FileName),
		log.WithLogLevel(logConf.Level),
		log.WithLogPath(logConf.LogPath),
		log.WithMaxSize(logConf.MaxSize),
		log.WithMaxBackups(logConf.MaxBackups))

	// 初始化DB
	gormcli.Init(
		gormcli.WithAddr(dbConf.Addr),
		gormcli.WithUser(dbConf.User),
		gormcli.WithPassword(dbConf.Password),
		gormcli.WithDataBase(dbConf.DataBase),
		gormcli.WithMaxIdleConn(dbConf.MaxIdleConn),
		gormcli.WithMaxOpenConn(dbConf.MaxOpenConn),
		gormcli.WithMaxIdleTime(dbConf.MaxIdleTime))

	cache.Init(
		cache.WithAddr(cacheConf.Addr),
		cache.WithPassWord(cacheConf.PassWord),
		cache.WithDB(cacheConf.DB),
		cache.WithPoolSize(cacheConf.PoolSize))

}

func main() {
	Init()
	router.InitRouterAndServe()
}
