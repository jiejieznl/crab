package impl

import (
	"context"
	"crab/modules/example2/internal/repo"
)

func newImplApi() repo.IApi {
	i := &implApi{baseDB: newBaseDB()}
	return i
}

type implApi struct {
	*baseDB
}

func (i *implApi) Find(ctx context.Context) {
	return
}

func (i *implApi) First(ctx context.Context) {
	return
}

func (i *implApi) Create(ctx context.Context) {
	return
}

func (i *implApi) Delete(ctx context.Context) {
	return
}

func (i *implApi) Update(ctx context.Context) {
	return
}
