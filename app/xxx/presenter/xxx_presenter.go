package presenter

import (
	"context"

	"xxx/app/domain/xxx"
)

type XxxPresenter interface {
	Create(ctx context.Context, dto *PresenterCreateDto, err error)
	Read(ctx context.Context, dto *PresenterReadDto, err error)
	ReadAll(ctx context.Context, dto *PresenterReadAllDto, err error)
	Update(ctx context.Context, dto *PresenterUpdateDto, err error)
	Delete(ctx context.Context, dto *PresenterDeleteDto, err error)
}

type PresenterCreateDto struct {
	Xxx xxx.Props
}

type PresenterReadDto struct {
	Xxx xxx.Props
}

type PresenterReadAllDto struct {
	Xxxs []xxx.Props
}

type PresenterUpdateDto struct {
	Xxx xxx.Props
}

type PresenterDeleteDto struct {
	Xxx xxx.Props
}
