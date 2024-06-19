package example

import "github.com/google/uuid"

type User struct {
	Model
	Id       uint      `json:"id" gorm:"primaryKey"`
	UUID     uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`    // 用户UUID
	Username string    `json:"username" gorm:"index;comment:用户登录名"` // 用户登录名
	Password string    `json:"-"  gorm:"comment:用户登录密码"`            // 用户登录密码
	Nickname string    `json:"nickname" gorm:"comment:用户昵称"`        // 用户昵称
	Avatar   string    `json:"avatar" gorm:"comment:用户头像"`          // 用户头像
	Status   uint      `json:"status" gorm:"comment:用户状态"`          //0禁用 1启用
}

func (User) TableName() string {
	return "sys_user"
}
