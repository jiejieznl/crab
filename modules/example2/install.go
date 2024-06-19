package example2

import (
	"crab/modules/example2/internal/controller"
	implRepo "crab/modules/example2/internal/repo/impl"
	implSrv "crab/modules/example2/internal/service/impl"
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
	implRepo.Initialize()
	implSrv.Initialize()
	controller.Initialize()
}
