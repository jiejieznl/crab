package example

import (
	"crab/modules/example/internal/controller"
	"crab/modules/example/internal/service"
	"github.com/gin-gonic/gin"
)

var (
	Module = &module{}
)

type module struct{}

func (m *module) RouterPrefix() string {
	return "api"
}

func (m *module) InstallRouter(group *gin.RouterGroup) {
	InstallRouter(group)
}

func (m *module) InitSequence() {
	service.Initialize()
	controller.Initialize()
}
