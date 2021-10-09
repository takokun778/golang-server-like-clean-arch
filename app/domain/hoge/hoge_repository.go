//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package hoge

import (
	"context"

	c "clean/app/domain/common"
)

type Repository interface {
	Save(ctx context.Context, model *Hoge) (*Hoge, error)
	Find(ctx context.Context, id c.Id) (*Hoge, error)
	FindAll(ctx context.Context) (*HogeList, error)
	Update(ctx context.Context, model *Hoge) (*Hoge, error)
	Delete(ctx context.Context, id c.Id) error
}
