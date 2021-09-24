package hoge

import (
	"context"

	c "clean/app/domain/common"
)

type Usecase interface {
	Create(ctx context.Context, input CreateUsecaseInput) (*CreateUsecaseOutput, *c.Error)
	Fetch(ctx context.Context, input FetchUsecaseInput) (*FetchUsecaseOutput, *c.Error)
	FetchAll(ctx context.Context, input FetchAllUsecaseInput) (*FetchAllUsecaseOutput, *c.Error)
	Update(ctx context.Context, input UpdateUsecaseInput) (*UpdateUsecaseOutput, *c.Error)
	Delete(ctx context.Context, input DeleteUsecaseInput) (*DeleteUsecaseOutput, *c.Error)
}

type CreateUsecaseInput struct {
	Name   Name
	Number Number
}

type CreateUsecaseOutput struct {
	Hoge *Hoge
}

type FetchUsecaseInput struct {
	Id c.Id
}

type FetchUsecaseOutput struct {
	Hoge *Hoge
}

type FetchAllUsecaseInput struct {
}

type FetchAllUsecaseOutput struct {
	HogeList *HogeList
}

type UpdateUsecaseInput struct {
	Id     c.Id
	Name   Name
	Number Number
}

type UpdateUsecaseOutput struct {
	Hoge *Hoge
}

type DeleteUsecaseInput struct {
	Id c.Id
}

type DeleteUsecaseOutput struct {
	Hoge *Hoge
}
