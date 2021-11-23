//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package xxx

import (
	"context"
)

type Repository interface {
	Save(ctx context.Context, input *RepositorySaveInput) (*RepositorySaveOutput, error)
	Find(ctx context.Context, input *RepositoryFindInput) (*RepositoryFindOutput, error)
	FindAll(ctx context.Context, input *RepositoryFindAllInput) (*RepositoryFindAllOutput, error)
	Update(ctx context.Context, input *RepositoryUpdateInput) (*RepositoryUpdateOutput, error)
	Delete(ctx context.Context, input *RepositoryDeleteInput) (*RepositoryDeleteOutput, error)
}

type RepositorySaveInput struct {
	Xxx Props
}

type RepositorySaveOutput struct {
	Xxx Props
}

type RepositoryFindInput struct {
	Id id
}

type RepositoryFindOutput struct {
	Xxx Props
}

type RepositoryFindAllInput struct {
}

type RepositoryFindAllOutput struct {
	Xxxs []Props
}

type RepositoryUpdateInput struct {
	Xxx Props
}

type RepositoryUpdateOutput struct {
	Xxx Props
}

type RepositoryDeleteInput struct {
	Id id
}

type RepositoryDeleteOutput struct {
}
