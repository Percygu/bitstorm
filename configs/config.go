package configs

import (
	"bytes"
	"fmt"
	rlog "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
)

const defaultTimeFormat = "2006-01-02 15:04:05"

// LogConf 日志配置
type LogConf struct {
	LogPattern string `yaml:"log_pattern" mapstructure:"log_pattern"` // 日志输出标准，终端输出/文件输出
	LogPath    string `yaml:"log_path" mapstructure:"log_path"`       // 日志路径
	SaveDays   uint   `yaml:"save_days" mapstructure:"save_days"`     // 日志保存天数
	Level      string `yaml:"level" mapstructure:"level"`             // 日志级别
}

// AppConf 服务配置
type AppConf struct {
	AppName string `yaml:"app_name" mapstructure:"app_name"` // 业务名
	Version string `yaml:"version" mapstructure:"version"`   // 版本
	Port    int    `yaml:"port" mapstructure:"port"`         // 端口
	RunMode string `yaml:"run_mode" mapstructure:"run_mode"` // 运行模式
}

// DbConf db配置结构
type DbConf struct {
	Host        string `yaml:"host" mapstructure:"host"`                   // db主机地址
	Port        string `yaml:"port" mapstructure:"port"`                   // db端口
	User        string `yaml:"user" mapstructure:"user"`                   // 用户名
	Password    string `yaml:"password" mapstructure:"password"`           // 密码
	DataBase    string `yaml:"database" mapstructure:"database"`           // db名
	MaxIdleConn int    `yaml:"max_idle_conn" mapstructure:"max_idle_conn"` // 最大空闲连接数
	MaxOpenConn int    `yaml:"max_open_conn" mapstructure:"max_open_conn"` // 最大打开的连接数
	MaxIdleTime int64  `yaml:"max_idle_time" mapstructure:"max_idle_time"` // 连接最大空闲时间
}

// RedisConf 配置
type RedisConf struct {
	Host     string `yaml:"rhost" mapstructure:"rhost"` // db主机地址
	Port     int    `yaml:"rport" mapstructure:"rport"` // db端口
	DB       int    `yaml:"rdb" mapstructure:"rdb"`
	PassWord string `yaml:"passwd" mapstructure:"passwd"`
	PoolSize int    `yaml:"poolsize" mapstructure:"poolsize"`
}

// GlobalConfig 业务配置结构体
type GlobalConfig struct {
	AppConfig   AppConf   `yaml:"app" mapstructure:"app"`
	LogConfig   LogConf   `yaml:"log" mapstructure:"log"`         // 日志配置
	DbConfig    DbConf    `yaml:"gormcli" mapstructure:"gormcli"` // db配置
	RedisConfig RedisConf `yaml:"redis" mapstructure:"redis"`     // redis配置
}

// logFormatter 日志格式化
type logFormatter struct {
	log.TextFormatter
}

var (
	config GlobalConfig // 全局业务配置文件
	once   sync.Once
)

// GetGlobalConfig 获取全局配置文件
func GetGlobalConfig() *GlobalConfig {
	once.Do(readConf)
	return &config
}

func readConf() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../configs")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("../../../../configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic("read config file err:" + err.Error())
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic("config file unmarshal err:" + err.Error())
	}
}

// InitConfig 初始化日志
func InitConfig() {
	globalConf := GetGlobalConfig()
	// 设置日志级别
	level, err := log.ParseLevel(globalConf.LogConfig.Level)
	if err != nil {
		panic("log level parse err:" + err.Error())
	}
	// 设置日志格式为json格式
	log.SetFormatter(&logFormatter{
		log.TextFormatter{
			DisableColors:   true,
			TimestampFormat: defaultTimeFormat,
			FullTimestamp:   true,
		}})
	log.SetReportCaller(true) // 打印文件位置，行号
	log.SetLevel(level)
	switch globalConf.LogConfig.LogPattern {
	case "stdout":
		log.SetOutput(os.Stdout)
	case "stderr":
		log.SetOutput(os.Stderr)
	case "file":
		logger, err := rlog.New(
			globalConf.LogConfig.LogPath+".%Y%m%d",
			// rlog.WithLinkName(globalConf.LogConf.LogPath),
			rlog.WithRotationCount(globalConf.LogConfig.SaveDays),
			// rlog.WithMaxAge(time.Minute*3),
			rlog.WithRotationTime(time.Hour*24),
		)
		if err != nil {
			panic("log conf err: " + err.Error())
		}
		log.SetOutput(logger)
	default:
		panic("log conf err, check log_pattern in logsvr.yaml")
	}
}

// Format 自定义日志输出格式
func (c *logFormatter) Format(entry *log.Entry) ([]byte, error) {
	prettyCaller := func(frame *runtime.Frame) string {
		_, fileName := filepath.Split(frame.File)
		return fmt.Sprintf("%s:%d", fileName, frame.Line)
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	b.WriteString(fmt.Sprintf("[%s] %s", entry.Time.Format(c.TimestampFormat), // 输出日志时间
		strings.ToUpper(entry.Level.String())))
	if entry.HasCaller() {
		b.WriteString(fmt.Sprintf("[%s]", prettyCaller(entry.Caller))) // 输出日志所在文件，行数位置
	}
	b.WriteString(fmt.Sprintf(" %s\n", entry.Message)) // 输出日志内容
	return b.Bytes(), nil
}
