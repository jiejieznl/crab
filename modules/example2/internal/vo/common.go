package vo

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

const (
	ERROR   = 400
	SUCCESS = 0
)

func result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func ResultOk(data any, msg string, c *gin.Context) {
	result(SUCCESS, data, msg, c)
}

func BindJSONOrRespondError(obj any, c *gin.Context) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		result(ERROR, nil, "传入参数错误", c)
		return false
	}
	return true
}

func ResultError(err error, c *gin.Context) {
	var e *Err
	if err != nil {
		if !errors.As(err, &e) {
			result(ERROR, nil, err.Error(), c)
			c.Abort()
			return
		}
		if e.code == 0 {
			e.code = ERROR
		}
		result(e.code, nil, e.msg, c)
		c.Abort()
		return
	}
	c.Abort()
	return
}
