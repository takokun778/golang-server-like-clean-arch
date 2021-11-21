//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package xxx

import (
	"context"

	"xxx/app/domain/common"
)

type Repository interface {
	Save(ctx context.Context, props Props) (Props, error)
	Find(ctx context.Context, id common.Id) (Props, error)
	FindAll(ctx context.Context) ([]Props, error)
	Update(ctx context.Context, props Props) (Props, error)
	Delete(ctx context.Context, id common.Id) error
}
