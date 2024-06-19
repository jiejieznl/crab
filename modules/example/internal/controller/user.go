package controller

import (
	"crab/modules/example/internal/service"
	"crab/modules/example/internal/util"
	"crab/modules/example/internal/vo"
	"github.com/gin-gonic/gin"
)

var User = new(_user)

func newUser() {
	User = &_user{srv: service.GetUser()}
}

type _user struct {
	srv *service.User
}

func (u *_user) Login(c *gin.Context) {
	var req vo.LoginReq
	if !vo.BindJSONOrRespondError(&req, c) {
		return
	}
	//第一种写法
	res, err := u.srv.Login(&req)
	if err != nil {
		vo.ResultError(err, c)
		return
	}
	//第二种 你也可以放弃initUser 就不用写initialize
	res, err = service.GetUser().Login(&req)
	if err != nil {
		vo.ResultError(err, c)
		return
	}
	vo.ResultOk(res, "登陆成功", c)
	return
}

func (u *_user) Auth(c *gin.Context) {
	vo.ResultOk(util.GetUserInfo(c), "jwt令牌信息", c)
}
