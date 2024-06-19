package config

import (
	"github.com/pelletier/go-toml/v2"
	"log"
	"os"
)

const (
	defaultConfigPath = "./config.toml"
)

var (
	conf *Config
)

func InitConfig(filePath string) {
	if len(filePath) == 0 {
		filePath = defaultConfigPath
	}
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("配置文件出错err:%v", err)
	}
	err = toml.Unmarshal(bytes, &conf)
	if err != nil {
		log.Fatalf("配置文件出错err:%v", err)
	}

}

func GetConfig() *Config {
	return conf
}

func GenDefaultConfig() {
	_conf := setDefaultConfig()

	// 将配置结构体编码为 TOML 格式的字节切片
	bytes, err := toml.Marshal(&_conf)
	if err != nil {
		log.Fatalf("生成默认配置文件出错err:%v", err)
	}

	// 创建 config_default.toml 文件
	file, err := os.Create("config_default.toml")
	if err != nil {
		log.Fatalf("创建文件出错err:%v", err)
	}
	defer file.Close()

	// 将 TOML 字节切片写入文件
	_, err = file.Write(bytes)
	if err != nil {
		log.Fatalf("写入配置到文件出错err:%v", err)
	}
}
