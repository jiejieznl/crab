package internal

import (
	"crab/pkg/config"
	"fmt"
)

func GenDefaultConfig() {
	config.GenDefaultConfig()
	fmt.Println("配置文件 config_default.toml 已生成")
}
