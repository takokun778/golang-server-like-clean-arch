package interactor

import (
	"context"
	"time"

	dErr "xxx/app/domain/error"
	"xxx/app/domain/xxx"
	"xxx/app/xxx/presenter"
)

func (i *xxxInteractor) ReadAll(ctx context.Context, dto *xxx.UsecaseReadAllDto, err error) {
	if err != nil {
		i.xxxPresenter.ReadAll(ctx, nil, err)
		return
	}

	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	foundAllXxx, err := i.xxxRepository.FindAll(timeOutCtx, &xxx.RepositoryFindAllItem{})

	if err != nil {
		i.xxxPresenter.ReadAll(timeOutCtx, nil, dErr.NewInternalServerError(err, ""))
		return
	}

	i.xxxPresenter.ReadAll(timeOutCtx, &presenter.PresenterReadAllDto{
		Xxxs: foundAllXxx,
	}, nil)
}
