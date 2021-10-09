//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package fuga

import (
	"context"

	"clean/app/domain/common"
)

type Usecase interface {
	Create(ctx context.Context, input UsecaseCreateInput) (*UsecaseCreateOutput, *common.Error)
	Fetch(ctx context.Context, input UsecaseFetchInput) (*UsecaseFetchOutput, *common.Error)
	FetchAll(ctx context.Context, input UsecaseFetchAllInput) (*UsecaseFetchAllOutput, *common.Error)
	Update(ctx context.Context, input UsecaseUpdateInput) (*UsecaseUpdateOutput, *common.Error)
	Delete(ctx context.Context, input UsecaseDeleteInput) (*UsecaseDeleteOutput, *common.Error)
}

type UsecaseCreateInput struct {
	Name   Name
	Number Number
}

type UsecaseCreateOutput struct {
	Fuga *Fuga
}

type UsecaseFetchInput struct {
	Id common.Id
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
	Id     common.Id
	Name   Name
	Number Number
}

type UsecaseUpdateOutput struct {
	Fuga *Fuga
}

type UsecaseDeleteInput struct {
	Id common.Id
}

type UsecaseDeleteOutput struct {
	Fuga *Fuga
}
