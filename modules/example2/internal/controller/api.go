package controller

import (
	"crab/modules/example2/internal/service"
	"crab/modules/example2/internal/vo"
	"github.com/gin-gonic/gin"
)

var Api = new(_api)

func newApi() {
	Api = &_api{srv: service.Api}
}

type _api struct {
	srv service.IApi
}

func (a *_api) Sum(c *gin.Context) {
	var req vo.ApiSumReq
	if !vo.BindJSONOrRespondError(&req, c) {
		return
	}
	res, err := a.srv.Sum(c.Copy(), &req)
	if err != nil {
		vo.ResultError(err, c)
		return
	}
	vo.ResultOk(res, "计算成功", c)
}
