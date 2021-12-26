package interactor

import (
	"context"
	"time"

	dErr "xxx/app/domain/error"
	"xxx/app/domain/xxx"
	"xxx/app/xxx/presenter"
)

func (i *xxxInteractor) Create(ctx context.Context, dto *xxx.UsecaseCreateDto) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	createdXxx := xxx.Create(dto.Name, dto.Number)

	item := &xxx.RepositorySaveItem{
		Xxx: createdXxx.Props(),
	}

	_, err := i.xxxRepository.Save(timeOutCtx, item)

	if err != nil {
		i.xxxPresenter.Create(ctx, nil, dErr.NewInternalServerError(err, ""))
		return
	}

	i.xxxPresenter.Create(ctx, &presenter.PresenterCreateDto{
		Xxx: createdXxx.Props(),
	}, nil)
}
