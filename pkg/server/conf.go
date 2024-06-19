package server

type Config struct {
	Mode string `toml:"mode" comment:"启动模式"`
	Port uint   `toml:"port" comment:"运行端口"`
}
