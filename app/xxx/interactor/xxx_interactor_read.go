package interactor

import (
	"context"
	"time"

	dErr "xxx/app/domain/error"
	"xxx/app/domain/xxx"
	"xxx/app/xxx/presenter"
)

func (i *xxxInteractor) Read(ctx context.Context, dto *xxx.UsecaseReadDto) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	foundXxx, err := i.xxxRepository.Find(timeOutCtx, &xxx.RepositoryFindItem{Id: dto.Id})

	if err != nil {
		i.xxxPresenter.Read(ctx, nil, dErr.NewInternalServerError(err, ""))
		return
	}

	i.xxxPresenter.Read(ctx, &presenter.PresenterReadDto{
		Xxx: foundXxx,
	}, nil)
}
