package logger

import (
	"ginServer/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var Log *zap.SugaredLogger
var logPath = "./log/"
var logLevel = config.Cfg.LogLevel
// 设置日志等级 debug< info < warn < error < panic < fatal ,默认为info

func init() {
	hook := lumberjack.Logger{
		Filename:   logPath + os.Args[0] + ".log", // 日志名
		MaxSize:    256,                           // 文件最大大小 M
		MaxAge:     30,                            // 文件 最多保存多少天
		MaxBackups: 2,                             // 最大备份数
		LocalTime:  false,
		Compress:   false,
	}

	encodeConfig := zapcore.EncoderConfig{
		MessageKey:    "TRACE",
		LevelKey:      "level",
		TimeKey:       "ts",
		CallerKey:     "file",
		FunctionKey:   "",
		StacktraceKey: "",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
		EncodeCaller:     zapcore.ShortCallerEncoder,
		EncodeName:       zapcore.FullNameEncoder,
		ConsoleSeparator: "",
	}

	var writes = []zapcore.WriteSyncer{zapcore.AddSync(&hook)}
	// 如果是debug模式，则同时打印文件和终端
	if logLevel == "debug" {
		writes = append(writes, zapcore.AddSync(os.Stdout))
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encodeConfig),
		zapcore.NewMultiWriteSyncer(writes...),
		getLogLevel(logLevel),
	)
	Log = zap.New(core, zap.AddCaller(), zap.Development()).Sugar()
}

// 设置日志等级
func getLogLevel(level string) zap.AtomicLevel {
	atomicLevel := zap.NewAtomicLevel()
	switch level {
	case "DEBUG":
		atomicLevel.SetLevel(zap.DebugLevel)
	case "INFO":
		atomicLevel.SetLevel(zap.InfoLevel)
	case "WARN":
		atomicLevel.SetLevel(zap.WarnLevel)
	case "ERROR":
		atomicLevel.SetLevel(zap.ErrorLevel)
	}
	return atomicLevel
}
