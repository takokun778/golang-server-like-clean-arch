//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package xxx

import (
	"context"
)

type UsecaseCreate interface {
	Handle(ctx context.Context, input *UsecaseCreateInput) (*UsecaseCreateOutput, error)
}

type UsecaseCreateInput struct {
	Name   name
	Number number
}

type UsecaseCreateOutput struct {
	Xxx Props
}

type UsecaseRead interface {
	Handle(ctx context.Context, input *UsecaseReadInput) (*UsecaseReadOutput, error)
}

type UsecaseReadInput struct {
	Id id
}

type UsecaseReadOutput struct {
	Xxx Props
}

type UsecaseReadAll interface {
	Handle(ctx context.Context, input *UsecaseReadAllInput) (*UsecaseReadAllOutput, error)
}

type UsecaseReadAllInput struct {
}

type UsecaseReadAllOutput struct {
	Xxxs []Props
}

type UsecaseUpdate interface {
	Handle(ctx context.Context, input *UsecaseUpdateInput) (*UsecaseUpdateOutput, error)
}

type UsecaseUpdateInput struct {
	Id     id
	Name   name
	Number number
}

type UsecaseUpdateOutput struct {
	Xxx Props
}

type UsecaseDelete interface {
	Handle(ctx context.Context, input *UsecaseDeleteInput) (*UsecaseDeleteOutput, error)
}

type UsecaseDeleteInput struct {
	Id id
}

type UsecaseDeleteOutput struct {
	Xxx Props
}
