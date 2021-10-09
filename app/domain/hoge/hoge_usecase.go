//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package hoge

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
	Hoge *Hoge
}

type UsecaseFetchInput struct {
	Id common.Id
}

type UsecaseFetchOutput struct {
	Hoge *Hoge
}

type UsecaseFetchAllInput struct {
}

type UsecaseFetchAllOutput struct {
	HogeList *HogeList
}

type UsecaseUpdateInput struct {
	Id     common.Id
	Name   Name
	Number Number
}

type UsecaseUpdateOutput struct {
	Hoge *Hoge
}

type UsecaseDeleteInput struct {
	Id common.Id
}

type UsecaseDeleteOutput struct {
	Hoge *Hoge
}
