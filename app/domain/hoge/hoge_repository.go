//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package hoge

import (
	"context"

	"clean/app/domain/common"
)

type Repository interface {
	Save(ctx context.Context, model Values) (Values, error)
	Find(ctx context.Context, id common.Id) (Values, error)
	FindAll(ctx context.Context) (ValuesList, error)
	Update(ctx context.Context, model Values) (Values, error)
	Delete(ctx context.Context, id common.Id) error
}
