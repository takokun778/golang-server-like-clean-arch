package fuga

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
	Fuga *Fuga
}

type FetchUsecaseInput struct {
	Id c.Id
}

type FetchUsecaseOutput struct {
	Fuga *Fuga
}

type FetchAllUsecaseInput struct {
}

type FetchAllUsecaseOutput struct {
	FugaList *FugaList
}

type UpdateUsecaseInput struct {
	Id     c.Id
	Name   Name
	Number Number
}

type UpdateUsecaseOutput struct {
	Fuga *Fuga
}

type DeleteUsecaseInput struct {
	Id c.Id
}

type DeleteUsecaseOutput struct {
	Fuga *Fuga
}
