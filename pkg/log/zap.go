package log

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZap(cfg *Config) *zap.Logger {
	cores, err := getCore(cfg)
	if err != nil {
		panic(fmt.Errorf("get cores error:%v", err))
	}
	ops := []zap.Option{zap.AddCaller(), zap.AddCallerSkip(1)}
	if cfg.EnableStacktrace {
		level, err := zapcore.ParseLevel(cfg.StacktraceLevel)
		if err != nil {
			level = zap.FatalLevel
		}
		ops = append(ops, zap.AddStacktrace(getMultiLevel(level)))
	}
	return zap.New(zapcore.NewTee(cores...), ops...)
}

func getCore(cfg *Config) ([]zapcore.Core, error) {
	cores := make([]zapcore.Core, 0, 6)
	level, err := zapcore.ParseLevel(cfg.Level)
	if err != nil {
		level = zap.InfoLevel
	}
	if cfg.EnableConsoleOut {
		cores = append(cores, getConsoleCore(getMultiLevel(level)))
	}
	if cfg.EnableFileOut {
		if cfg.EnableMixedSave { // 开启混合保存，true:多等级日志保存一个文件，false:不同等级分文件存储
			ws, err := createWriter(cfg, "2006-01-02-1504")
			if err != nil {
				return nil, err
			}
			cores = append(cores, getFileCore(ws, getMultiLevel(level)))
		} else {
			for i := level; i <= zapcore.FatalLevel; i++ {
				if i == zapcore.DPanicLevel || i == zapcore.PanicLevel {
					continue
				}
				ws, err := createWriter(cfg, fmt.Sprintf("2006-01-02-1504.%s", i.String()))
				if err != nil {
					return nil, err
				}
				cores = append(cores, getFileCore(ws, getEqualLevel(i)))
			}
		}
	}
	return cores, nil
}

func getConsoleCore(level zap.LevelEnablerFunc) zapcore.Core {
	return zapcore.NewCore(zapcore.NewConsoleEncoder(getEncoderConfig(true)), os.Stdout, level)
}

func getFileCore(file zapcore.WriteSyncer, level zap.LevelEnablerFunc) zapcore.Core {
	return zapcore.NewCore(zapcore.NewConsoleEncoder(getEncoderConfig(false)), file, level)
}

func getMultiLevel(l zapcore.Level) zap.LevelEnablerFunc {
	return func(level zapcore.Level) bool {
		return level >= l
	}
}

func getEqualLevel(l zapcore.Level) zap.LevelEnablerFunc {
	return func(level zapcore.Level) bool {
		return level == l
	}
}

func getEncoderConfig(color bool) zapcore.EncoderConfig {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "name",
		CallerKey:     "call",
		StacktraceKey: "stack",
		LineEnding:    zapcore.DefaultLineEnding,
		//EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	if color {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}
	return encoderConfig
}

func customTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}
