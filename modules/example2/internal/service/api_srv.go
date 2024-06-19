package service

import (
	"context"
	"crab/modules/example2/internal/vo"
)

type IApi interface {
	Sum(ctx context.Context, req *vo.ApiSumReq) (*vo.ApiSumRes, error)
}

var Api IApi
