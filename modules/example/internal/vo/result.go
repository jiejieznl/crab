package vo

import "fmt"

type Err struct {
	code int
	msg  string
}

func NewErr() *Err {
	return &Err{}
}

func (e *Err) Error() string {
	return fmt.Sprintf("ErrCode:%d ErrMsg:%s", e.code, e.msg)
}

func (e *Err) Code(code int) *Err {
	e.code = code
	return e
}

func (e *Err) Msg(msg string) *Err {
	e.msg = msg
	return e
}
