//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package fuga

import (
	"clean/app/domain/common"
	"context"
)

type Repository interface {
	Save(ctx context.Context, values Values) (Values, error)
	Find(ctx context.Context, id common.Id) (Values, error)
	FindAll(ctx context.Context) (ValuesList, error)
	Update(ctx context.Context, values Values) (Values, error)
	Delete(ctx context.Context, id common.Id) error
}
