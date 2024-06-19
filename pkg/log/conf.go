package log

type Config struct {
	// Level: Debug Info Warn Error Fatal
	Level            string `toml:"level" comment:"日志等级"`
	Directory        string `toml:"directory" comment:"日志保存路径"`
	MaxSize          int    `toml:"max-size" comment:"日志单文件最大大小(MB)"`
	MaxAge           int    `toml:"max-age" comment:"日志文件过期时间(天)"`
	StacktraceLevel  string `toml:"stacktrace-level" comment:"开启栈捕获"`
	EnableStacktrace bool   `toml:"enable-stacktrace" comment:"启用堆栈跟踪"`
	EnableFileOut    bool   `toml:"enable-file-out" comment:"启用文件输出"`
	EnableMixedSave  bool   `toml:"enable-mixed-save" comment:"启用混合保存"`
	EnableConsoleOut bool   `toml:"enable-console-out" comment:"启用控制台输出"`
}

func setDefaultConfig() *Config {
	return &Config{
		Level:            "info",
		Directory:        "log",
		MaxSize:          1024,
		MaxAge:           0,
		StacktraceLevel:  "fatal",
		EnableStacktrace: false,
		EnableFileOut:    false,
		EnableMixedSave:  true,
		EnableConsoleOut: true,
	}
}
