package log

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var (
	log *zap.Logger
)

// 默认日志组件使用zap
type Logger interface {
	Error(v ...interface{})
	Warn(v ...interface{})
	Info(v ...interface{})
	Debug(v ...interface{})
	Errorf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Debugf(format string, v ...interface{})
}

var (
	logger Logger
)

func Init(opts ...Option) {
	// 初始化
	logger = newSugarLogger(NewOptions(opts...))
}

// Options 选项配置
type Options struct {
	LogPath    string // 日志路径
	LogName    string // 日志名称
	LogLevel   string // 日志级别
	FileName   string // 文件名称
	MaxAge     int    // 日志保留时间，以天为单位
	MaxSize    int    // 日志保留大小，以 M 为单位
	MaxBackups int    // 保留文件个数
	Compress   bool   // 是否压缩
}

// Option 选项方法
type Option func(*Options)

// NewOptions 初始化
func NewOptions(opts ...Option) Options {
	// 默认配置
	options := Options{
		LogName:    "shortUrlX",
		LogLevel:   "info",
		FileName:   "app.log",
		MaxAge:     10,
		MaxSize:    100,
		MaxBackups: 3,
		Compress:   true,
	}
	for _, opt := range opts {
		opt(&options)
	}
	return options
}

// WithLogLevel 日志级别
func WithLogLevel(level string) Option {
	return func(o *Options) {
		o.LogLevel = level
	}
}

// WithFileName 日志文件
func WithFileName(filename string) Option {
	return func(o *Options) {
		o.FileName = filename
	}
}
func WithLogPath(logPath string) Option {
	return func(o *Options) {
		o.LogPath = logPath
	}
}
func WithMaxAge(maxAge int) Option {
	return func(o *Options) {
		o.MaxAge = maxAge
	}
}
func WithMaxSize(maxSize int) Option {
	return func(o *Options) {
		o.MaxSize = maxSize
	}
}
func WithMaxBackups(maxBackups int) Option {
	return func(o *Options) {
		o.MaxBackups = maxBackups
	}
}
func WithCompress(compress bool) Option {
	return func(o *Options) {
		o.Compress = compress
	}
}

type zapLoggerWrapper struct {
	*zap.SugaredLogger
	options Options
}

func newSugarLogger(options Options) *zapLoggerWrapper {
	w := &zapLoggerWrapper{options: options}
	encoder := w.getEncoder()
	w.setSugaredLogger(encoder)
	return w
}
func (w *zapLoggerWrapper) setSugaredLogger(encoder zapcore.Encoder) {
	var coreArr []zapcore.Core
	// info文件writeSyncer
	// 日志级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { // error级别
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { // info和debug级别,debug级别是最低的
		if w.options.LogLevel == "debug" {
			return lev < zap.ErrorLevel && lev >= zap.DebugLevel
		} else {
			return lev < zap.ErrorLevel && lev >= zap.InfoLevel
		}
	})
	infoFileWriteSyncer := w.getLogWriter("info_")
	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowPriority)
	errorFileWriteSyncer := w.getLogWriter("error_")
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority)
	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)
	log = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller(), zap.AddCallerSkip(1)) // zap.AddCaller()为显示文件名和行号，可省略
	w.SugaredLogger = log.Sugar()
}
func (w *zapLoggerWrapper) getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// NewConsoleEncoder 打印更符合人们观察的方式
	return zapcore.NewConsoleEncoder(encoderConfig)
}
func (w *zapLoggerWrapper) getLogWriter(typeName string) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   w.options.LogPath + typeName + w.options.FileName, // 日志文件存放目录，
		MaxSize:    w.options.MaxSize,                                 // 文件大小限制,单位MB
		MaxBackups: w.options.MaxBackups,                              // 最大保留日志文件数量
		MaxAge:     w.options.MaxAge,                                  // 日志文件保留天数
		Compress:   w.options.Compress,                                // 是否压缩处理
	})
}

// GetDefaultLogger 获取默认日志实现
func GetDefaultLogger() Logger {
	return logger
}

// Debugf 打印 Debug 日志
func Debugf(format string, args ...interface{}) {
	GetDefaultLogger().Debugf(format, args...)
}

// Infof 打印 Info 日志
func Infof(format string, args ...interface{}) {
	GetDefaultLogger().Infof(format, args...)
}

// Warnf 打印 Warn 日志
func Warnf(format string, args ...interface{}) {
	GetDefaultLogger().Warnf(format, args...)
}

// Errorf 打印 Error 日志
func Errorf(format string, args ...interface{}) {
	GetDefaultLogger().Errorf(format, args...)
}

// DebugContext 打印 Debug 日志
func DebugContext(ctx context.Context, args ...interface{}) {
	GetDefaultLogger().Debug(args...)
}

// DebugContextf 打印 Debug 日志
func DebugContextf(ctx context.Context, format string, args ...interface{}) {
	GetDefaultLogger().Debugf(format, args...)
}

// InfoContext 打印 Info 日志
func InfoContext(ctx context.Context, args ...interface{}) {
	GetDefaultLogger().Info(args...)
}

// InfoContextf 打印 Info 日志
func InfoContextf(ctx context.Context, format string, args ...interface{}) {
	value := ctx.Value("req_id")
	args = append([]interface{}{value}, args...)
	GetDefaultLogger().Infof("req_id:%s, "+format, args...)
}

// WarnContext 打印 Warn 日志
func WarnContext(ctx context.Context, args ...interface{}) {
	GetDefaultLogger().Warn(args...)
}

// WarnContextf 打印 Warn 日志
func WarnContextf(ctx context.Context, format string, args ...interface{}) {
	GetDefaultLogger().Warnf(format, args...)
}

// ErrorContext 打印 Error 日志
func ErrorContext(ctx context.Context, args ...interface{}) {
	GetDefaultLogger().Error(args...)
}
func ErrorContextf(ctx context.Context, format string, args ...interface{}) {
	GetDefaultLogger().Errorf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	Errorf(format, args...)
}
