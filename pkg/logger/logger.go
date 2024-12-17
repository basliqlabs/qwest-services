package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	Level      string `koanf:"level"`
	Encoding   string `koanf:"encoding"`
	OutputPath string `koanf:"output_path"`
	MaxSize    int    `koanf:"max_size"`
	MaxBackups int    `koanf:"max_backups"`
	MaxAge     int    `koanf:"max_age"`
	Compress   bool   `koanf:"compress"`
}

var globalLogger *zap.Logger

func Init(cfg Config, env string) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
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

	level, err := zapcore.ParseLevel(cfg.Level)
	if err != nil {
		fmt.Printf("invalid log level %s, defaulting to info: %v\n", cfg.Level, err)
		level = zapcore.InfoLevel
	}

	var encoder zapcore.Encoder
	if cfg.Encoding == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	fileSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.OutputPath,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	})

	var core zapcore.Core
	if env == "development" {
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level),
			zapcore.NewCore(encoder, fileSyncer, level),
		)
	} else {
		core = zapcore.NewCore(encoder, fileSyncer, level)
	}

	globalLogger = zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
}

func L() *zap.Logger {
	if globalLogger == nil {
		panic("logger not initialized")
	}
	return globalLogger
}
