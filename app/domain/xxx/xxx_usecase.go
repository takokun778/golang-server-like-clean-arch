//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package xxx

import (
	"context"
)

type Usecase interface {
	Create(ctx context.Context, input *UsecaseCreateInput) (*UsecaseCreateOutput, error)
	Read(ctx context.Context, input *UsecaseReadInput) (*UsecaseReadOutput, error)
	ReadAll(ctx context.Context, input *UsecaseReadAllInput) (*UsecaseReadAllOutput, error)
	Update(ctx context.Context, input *UsecaseUpdateInput) (*UsecaseUpdateOutput, error)
	Delete(ctx context.Context, input *UsecaseDeleteInput) (*UsecaseDeleteOutput, error)
}

type UsecaseCreateInput struct {
	Name   name
	Number number
}

type UsecaseCreateOutput struct {
	Xxx Props
}

type UsecaseReadInput struct {
	Id id
}

type UsecaseReadOutput struct {
	Xxx Props
}

type UsecaseReadAllInput struct {
}

type UsecaseReadAllOutput struct {
	Xxxs []Props
}

type UsecaseUpdateInput struct {
	Id     id
	Name   name
	Number number
}

type UsecaseUpdateOutput struct {
	Xxx Props
}

type UsecaseDeleteInput struct {
	Id id
}

type UsecaseDeleteOutput struct {
	Xxx Props
}
