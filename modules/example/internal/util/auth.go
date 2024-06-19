package util

import (
	"crab/modules/example/internal/vo"
	"errors"
	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) (string, error) {
	token := c.GetHeader("Authorization")
	if len(token) == 0 {
		return "", errors.New("没有传入Token")
	}
	return token, nil
}

func GetUserInfo(c *gin.Context) *vo.UserInfo {
	value, _ := c.Get("user")
	return value.(*vo.UserInfo)
}
