//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package fuga

import (
	"context"

	"clean/app/domain/common"
)

type Usecase interface {
	Create(ctx context.Context, input UsecaseCreateInput) (UsecaseCreateOutput, error)
	Fetch(ctx context.Context, input UsecaseFetchInput) (UsecaseFetchOutput, error)
	FetchAll(ctx context.Context, input UsecaseFetchAllInput) (UsecaseFetchAllOutput, error)
	Update(ctx context.Context, input UsecaseUpdateInput) (UsecaseUpdateOutput, error)
	Delete(ctx context.Context, input UsecaseDeleteInput) (UsecaseDeleteOutput, error)
}

type UsecaseCreateInput struct {
	Name   Name
	Number Number
}

type UsecaseCreateOutput struct {
	Fuga Values
}

type UsecaseFetchInput struct {
	Id common.Id
}

type UsecaseFetchOutput struct {
	Fuga Values
}

type UsecaseFetchAllInput struct {
}

type UsecaseFetchAllOutput struct {
	FugaList ValuesList
}

type UsecaseUpdateInput struct {
	Id     common.Id
	Name   Name
	Number Number
}

type UsecaseUpdateOutput struct {
	Fuga Values
}

type UsecaseDeleteInput struct {
	Id common.Id
}

type UsecaseDeleteOutput struct {
	Fuga Values
}
