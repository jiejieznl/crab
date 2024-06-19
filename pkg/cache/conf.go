package cache

type Config struct {
	Source string `toml:"source" comment:"缓存类型"`
	Redis  Redis  `toml:"redis" comment:"redis配置"`
	Local  Local  `toml:"local" comment:"本地缓存配置"`
}

type Redis struct {
	Addr         string `toml:"addr" comment:"redis运行地址"`
	Port         uint   `toml:"port" comment:"redis运行端口"`
	Password     string `toml:"password" comment:"redis密码"`
	DB           int    `toml:"sql" comment:"redis数据库"`
	DialTimeout  uint   `toml:"dial-timeout" comment:"连接到redis服务器的超时时间(秒)"`
	ReadTimeout  uint   `toml:"read-timeout" comment:"读取操作的超时时间(秒)"`
	WriteTimeout uint   `toml:"write-timeout" comment:"写入操作的超时时间(秒)"`
}

type Local struct {
	DefaultExpiration int `toml:"defaultExpiration" comment:"默认过期时间(分钟)"`
	CleanupInterval   int `toml:"cleanupInterval" comment:"定时清理间隔(分钟)"`
}
