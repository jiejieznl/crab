package service

import (
	"crab/model/example"
	"crab/modules/example/internal/vo"
	"crab/pkg/jwt"
	"crab/pkg/log"
	"crab/pkg/sql"
	"errors"
	"gorm.io/gorm"
	"time"
)

var _user *User

func newUser() {
	_user = &User{
		db:  sql.GetDB(),
		log: log.GetLog(),
	}
}

func GetUser() *User {
	if _user == nil {
		panic("service User not impl")
	}
	return _user
}

type User struct {
	db  *gorm.DB
	log log.ILog
}

func (u *User) Login(req *vo.LoginReq) (*vo.LoginRes, error) {
	var user example.User
	if err := u.db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return nil, errors.New("帐号不存在或密码错误")
	}

	if user.Password != req.Password {
		return nil, errors.New("密码错误")
	}

	if user.Status != 1 {
		return nil, errors.New("帐号已被禁用")
	}

	info := vo.UserInfo{
		Id:   user.Id,
		UUID: user.UUID,
	}
	info.IssuedAt = time.Now()
	info.ExpirationTime, _ = jwt.GetAccessExpTime(info.GetIssuedAt())
	token, err := jwt.GenerationToken(&info)
	if err != nil {
		return nil, errors.New("token下发失败")
	}
	return &vo.LoginRes{Token: token}, nil
}
