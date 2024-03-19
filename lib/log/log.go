package log

import (
	"fmt"
	"os"
	"sync"
	"viry_sun/lib/config"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var L *zap.Logger

var once sync.Once

func init() {
	once.Do(func() {
		initLog()

		//配置变更时重新初始化
		config.RegisterEvent("log", func() {
			initLog()
		})
	})
}

func initLog() {
	// 定义处理日志级别
	allPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.Level(config.C.Stdout.MinLevel) &&
			lvl <= zapcore.Level(config.C.Stdout.MaxLevel)
	})
	var highPriority zap.LevelEnablerFunc
	var lowPriority zap.LevelEnablerFunc

	//日志文件切割配置
	var errorWriter zapcore.WriteSyncer
	var logWriter zapcore.WriteSyncer

	var core zapcore.Core

	//记录方式
	allWriter := zapcore.AddSync(os.Stdout) //打印到控制台

	//记录格式
	//内置格式 {"level":"debug","ts":1572160754.994731,"msg":"xxxx"}
	//encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	//自定义格式
	encoder := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	pCores := []zapcore.Core{zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), allWriter, allPriority)}

	if config.C.Error.File != "" {
		// 定义处理日志级别
		highPriority = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.Level(config.C.Error.MinLevel) &&
				lvl <= zapcore.Level(config.C.Error.MaxLevel)
		})
		//日志文件切割配置
		errorWriter = zapcore.AddSync(&lumberjack.Logger{
			Filename:   config.C.Error.File,
			MaxSize:    config.C.Error.MaxSize,    //最大M数，超过则切割
			MaxBackups: config.C.Error.MaxBackups, //最大文件保留数，超过就删除最老的日志文件
			MaxAge:     config.C.Error.MaxAge,     //保存N天
			Compress:   false,                     //是否压缩
		})
		pCores = append(pCores, zapcore.NewCore(zapcore.NewJSONEncoder(encoder), errorWriter, highPriority))
	}

	fmt.Printf("log file %v", config.C.Log.File)
	if config.C.Log.File != "" {
		// 定义处理日志级别
		lowPriority = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.Level(config.C.Log.MinLevel) &&
				lvl <= zapcore.Level(config.C.Log.MaxLevel)
		})
		//日志文件切割配置
		logWriter = zapcore.AddSync(&lumberjack.Logger{
			Filename:   config.C.Log.File,
			MaxSize:    config.C.Log.MaxSize,    //最大M数，超过则切割
			MaxBackups: config.C.Log.MaxBackups, //最大文件保留数，超过就删除最老的日志文件
			MaxAge:     config.C.Log.MaxAge,     //保存N天
			Compress:   false,                   //是否压缩
		})
		pCores = append(pCores, zapcore.NewCore(zapcore.NewJSONEncoder(encoder), logWriter, lowPriority))
	}

	// zapcore.Cores, then tee the four cores together.
	core = zapcore.NewTee(pCores...)

	// From a zapcore.Core, it's easy to construct a Logger.
	L = zap.New(core)
	defer L.Sync()
}
