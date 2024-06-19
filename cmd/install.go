package main

import (
	"crab/modules/example"
	"crab/modules/example2"
	"github.com/gin-gonic/gin"
)

type IModule interface {
	InitSequence()
	InstallRouter(group *gin.RouterGroup)
	RouterPrefix() string
}

var (
	modules = []IModule{
		example.Module,
		example2.Module,
	}
)

func moduleInstall(app *gin.Engine) {
	for _, module := range modules {
		module.InitSequence()
		module.InstallRouter(app.Group(module.RouterPrefix()))
	}
}
