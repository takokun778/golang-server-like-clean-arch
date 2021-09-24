package hoge

import (
	"context"

	c "clean/app/domain/common"
)

type Repository interface {
	Save(ctx context.Context, hoge *Hoge) (*Hoge, error)
	Find(ctx context.Context, id c.Id) (*Hoge, error)
	FindAll(ctx context.Context) (*HogeList, error)
	Update(ctx context.Context, hoge *Hoge) (*Hoge, error)
	Delete(ctx context.Context, id c.Id) error
}
