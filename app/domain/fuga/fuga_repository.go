//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package fuga

import (
	c "clean/app/domain/common"
	"context"
)

type Repository interface {
	Save(ctx context.Context, model *Fuga) (*Fuga, error)
	Find(ctx context.Context, id c.Id) (*Fuga, error)
	FindAll(ctx context.Context) (*FugaList, error)
	Update(ctx context.Context, model *Fuga) (*Fuga, error)
	Delete(ctx context.Context, id c.Id) error
}
