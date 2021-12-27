//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package xxx

import (
	"context"
)

type Usecase interface {
	Create(ctx context.Context, dto *UsecaseCreateDto, err error)
	Read(ctx context.Context, dto *UsecaseReadDto, err error)
	ReadAll(ctx context.Context, dto *UsecaseReadAllDto, err error)
	Update(ctx context.Context, dto *UsecaseUpdateDto, err error)
	Delete(ctx context.Context, dto *UsecaseDeleteDto, err error)
}

type UsecaseCreateDto struct {
	Name   name
	Number number
}

type UsecaseReadDto struct {
	Id id
}

type UsecaseReadAllDto struct {
}

type UsecaseUpdateDto struct {
	Id     id
	Name   name
	Number number
}

type UsecaseDeleteDto struct {
	Id id
}
