//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package xxx

import (
	"context"
)

type Usecase interface {
	Create(ctx context.Context, dto *UsecaseCreateDto)
	Read(ctx context.Context, dto *UsecaseReadDto)
	ReadAll(ctx context.Context, dto *UsecaseReadAllDto)
	Update(ctx context.Context, dto *UsecaseUpdateDto)
	Delete(ctx context.Context, dto *UsecaseDeleteDto)
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
