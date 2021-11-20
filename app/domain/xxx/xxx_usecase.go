//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package xxx

import (
	"context"

	"xxx/app/domain/common"
)

type Usecase interface {
	Create(ctx context.Context, input *UsecaseCreateInput) (*UsecaseCreateOutput, error)
	Read(ctx context.Context, input *UsecaseReadInput) (*UsecaseReadOutput, error)
	ReadAll(ctx context.Context, input *UsecaseReadAllInput) (*UsecaseReadAllOutput, error)
	Update(ctx context.Context, input *UsecaseUpdateInput) (*UsecaseUpdateOutput, error)
	Delete(ctx context.Context, input *UsecaseDeleteInput) (*UsecaseDeleteOutput, error)
}

type UsecaseCreateInput struct {
	Name   Name
	Number Number
}

type UsecaseCreateOutput struct {
	Xxx Values
}

type UsecaseReadInput struct {
	Id common.Id
}

type UsecaseReadOutput struct {
	Xxx Values
}

type UsecaseReadAllInput struct {
}

type UsecaseReadAllOutput struct {
	Xxxs []Values
}

type UsecaseUpdateInput struct {
	Id     common.Id
	Name   Name
	Number Number
}

type UsecaseUpdateOutput struct {
	Xxx Values
}

type UsecaseDeleteInput struct {
	Id common.Id
}

type UsecaseDeleteOutput struct {
	Xxx Values
}
