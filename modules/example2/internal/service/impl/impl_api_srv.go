package impl

import (
	"context"
	"crab/modules/example2/internal/repo"
	"crab/modules/example2/internal/service"
	"crab/modules/example2/internal/vo"
	"crab/pkg/log"
)

func newImplApi() service.IApi {
	return &implApi{
		log: log.GetLog(),
		ra:  repo.Api,
	}
}

type implApi struct {
	log log.ILog
	ra  repo.IApi
}

func (i *implApi) Sum(ctx context.Context, req *vo.ApiSumReq) (*vo.ApiSumRes, error) {
	i.ra.Find(ctx)
	i.ra.First(ctx)
	i.ra.Create(ctx)
	i.ra.Delete(ctx)
	i.ra.Update(ctx)
	return &vo.ApiSumRes{
		Sum: req.Num1 + req.Num2,
	}, nil
}
