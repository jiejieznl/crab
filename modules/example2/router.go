package example2

import (
	"crab/modules/example2/internal/controller"
	"github.com/gin-gonic/gin"
)

func InstallRouter(group *gin.RouterGroup) {
	api := controller.Api
	group.POST("sum", api.Sum)
}
