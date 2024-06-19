package middleware

import (
	"crab/modules/example/internal/util"
	"crab/modules/example/internal/vo"
	"crab/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := util.GetToken(c)
		if err != nil {
			c.AbortWithStatusJSON(200, vo.Response{
				Code:    401,
				Data:    nil,
				Message: err.Error(),
			})
			return
		}
		claims, err := jwt.VerifyToken[*vo.UserInfo](token)
		if err != nil {
			c.AbortWithStatusJSON(200, vo.Response{
				Code:    401,
				Data:    nil,
				Message: err.Error(),
			})
			return
		}
		c.Set("user", claims)
		c.Next()
	}
}
