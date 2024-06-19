package repo

import "context"

type IApi interface {
	Find(ctx context.Context)
	First(ctx context.Context)
	Create(ctx context.Context)
	Delete(ctx context.Context)
	Update(ctx context.Context)
}

var Api IApi
