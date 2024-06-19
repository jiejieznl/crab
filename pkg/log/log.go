package log

import (
	"go.uber.org/zap"
)

type ILog interface {
	NameSpace(name string) ILog
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)
}

var (
	runtimeILog ILog
	conf        *Config
)

func Register(cfg *Config) {
	conf = cfg
	if cfg.Level == "" {
		conf = setDefaultConfig()
	}
	runtimeILog = &logImpl{zap: newZap(conf).Sugar()}
}

func GetLog() ILog {
	if runtimeILog == nil {
		panic("log not init")
	}
	return runtimeILog
}

type logImpl struct {
	zap *zap.SugaredLogger
}

func (l *logImpl) NameSpace(name string) ILog        { return &logImpl{zap: l.zap.Named(name)} }
func (l *logImpl) Debugf(format string, args ...any) { l.zap.Debugf(format, args...) }
func (l *logImpl) Infof(format string, args ...any)  { l.zap.Infof(format, args...) }
func (l *logImpl) Warnf(format string, args ...any)  { l.zap.Warnf(format, args...) }
func (l *logImpl) Errorf(format string, args ...any) { l.zap.Errorf(format, args...) }
func (l *logImpl) Fatalf(format string, args ...any) { l.zap.Fatalf(format, args...) }
func (l *logImpl) Debug(args ...any)                 { l.zap.Debug(args...) }
func (l *logImpl) Info(args ...any)                  { l.zap.Info(args...) }
func (l *logImpl) Warn(args ...any)                  { l.zap.Warn(args...) }
func (l *logImpl) Error(args ...any)                 { l.zap.Error(args...) }
func (l *logImpl) Fatal(args ...any)                 { l.zap.Fatal(args...) }
