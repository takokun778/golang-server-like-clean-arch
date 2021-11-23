//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package xxx

import (
	"context"
)

type Repository interface {
	Save(ctx context.Context, item *RepositorySaveItem) (Props, error)
	Find(ctx context.Context, item *RepositoryFindItem) (Props, error)
	FindAll(ctx context.Context, item *RepositoryFindAllItem) ([]Props, error)
	Update(ctx context.Context, item *RepositoryUpdateItem) (Props, error)
	Delete(ctx context.Context, item *RepositoryDeleteItem) error
}

type RepositorySaveItem struct {
	Xxx Props
}

type RepositoryFindItem struct {
	Id id
}

type RepositoryFindAllItem struct {
}

type RepositoryUpdateItem struct {
	Xxx Props
}

type RepositoryDeleteItem struct {
	Id id
}
