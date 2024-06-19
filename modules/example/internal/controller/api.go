package controller

import (
	"crab/modules/example/internal/vo"
	"github.com/gin-gonic/gin"
)

var Api = new(api)

type api struct {
}

func (a *api) Ping(c *gin.Context) {
	var req vo.LoginReq
	if !vo.BindJSONOrRespondError(&req, c) {
		return
	}
	vo.ResultOk(nil, "ok", c)
	return
}
