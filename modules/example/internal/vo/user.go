package vo

import (
	"github.com/google/uuid"
	"time"
)

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRes struct {
	Token string `json:"token"`
}

type UserInfo struct {
	Id             uint      `json:"id"`
	UUID           uuid.UUID `json:"uuid"`
	IssuedAt       time.Time `json:"issued_at"`
	ExpirationTime time.Time `json:"expiration_time"`
}

func (u *UserInfo) GetSubject() string           { return u.UUID.String() }
func (u *UserInfo) GetIssuer() string            { return "SYSTEM" }
func (u *UserInfo) GetIssuedAt() time.Time       { return u.IssuedAt }
func (u *UserInfo) GetExpirationTime() time.Time { return u.ExpirationTime }
