package jwt

type Config struct {
	//SigningKey 可以删除此配置 让程序每次运行的时候随机UUID 那么程序重启会导致所有用户都要重新登陆
	SigningKey string `toml:"signing-key" comment:"jwt加密key,请确保您的系统key都不一致"`
}
