//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package xxx

import (
	"context"

	"xxx/app/domain/common"
)

type Repository interface {
	Save(ctx context.Context, model Values) (Values, error)
	Find(ctx context.Context, id common.Id) (Values, error)
	FindAll(ctx context.Context) ([]Values, error)
	Update(ctx context.Context, model Values) (Values, error)
	Delete(ctx context.Context, id common.Id) error
}
