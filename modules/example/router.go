package example

import (
	"crab/modules/example/internal/controller"
	"crab/modules/example/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InstallRouter(group *gin.RouterGroup) {
	common := group.Group("/")
	{
		api := controller.Api
		user := controller.User
		common.GET("ping", api.Ping)
		common.POST("login", user.Login)
	}

	auth := group.Group("/auth")
	{
		auth.Use(middleware.AuthJwt())
		user := controller.User
		auth.GET("", user.Auth)
	}
}
