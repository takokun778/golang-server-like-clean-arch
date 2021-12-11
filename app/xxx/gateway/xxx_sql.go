package gateway

import (
	"context"

	"github.com/google/uuid"
)

type XxxSql interface {
	InsertXxx(ctx context.Context, xxx XxxEntity) error
	SelectXxx(ctx context.Context, id uuid.UUID) (XxxEntity, error)
	SelectXxxAll(ctx context.Context) ([]XxxEntity, error)
	UpdateXxx(ctx context.Context, xxx XxxEntity) error
	DeleteXxx(ctx context.Context, id uuid.UUID) error
}
