//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package fuga

import (
	"context"

	c "clean/app/domain/common"
)

type Usecase interface {
	Create(ctx context.Context, input UsecaseCreateInput) (*UsecaseCreateOutput, *c.Error)
	Fetch(ctx context.Context, input UsecaseFetchInput) (*UsecaseFetchOutput, *c.Error)
	FetchAll(ctx context.Context, input UsecaseFetchAllInput) (*UsecaseFetchAllOutput, *c.Error)
	Update(ctx context.Context, input UsecaseUpdateInput) (*UsecaseUpdateOutput, *c.Error)
	Delete(ctx context.Context, input UsecaseDeleteInput) (*UsecaseDeleteOutput, *c.Error)
}

type UsecaseCreateInput struct {
	Name   Name
	Number Number
}

type UsecaseCreateOutput struct {
	Fuga *Fuga
}

type UsecaseFetchInput struct {
	Id c.Id
}

type UsecaseFetchOutput struct {
	Fuga *Fuga
}

type UsecaseFetchAllInput struct {
}

type UsecaseFetchAllOutput struct {
	FugaList *FugaList
}

type UsecaseUpdateInput struct {
	Id     c.Id
	Name   Name
	Number Number
}

type UsecaseUpdateOutput struct {
	Fuga *Fuga
}

type UsecaseDeleteInput struct {
	Id c.Id
}

type UsecaseDeleteOutput struct {
	Fuga *Fuga
}
