package interactor

import (
	"context"
	"time"

	dErr "xxx/app/domain/error"
	"xxx/app/domain/xxx"
	"xxx/app/xxx/presenter"
)

func (i *xxxInteractor) ReadAll(ctx context.Context, dto *xxx.UsecaseReadAllDto) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	foundAllXxx, err := i.xxxRepository.FindAll(timeOutCtx, &xxx.RepositoryFindAllItem{})

	if err != nil {
		i.xxxPresenter.ReadAll(ctx, nil, dErr.NewInternalServerError(err, ""))
		return
	}

	i.xxxPresenter.ReadAll(ctx, &presenter.PresenterReadAllDto{
		Xxxs: foundAllXxx,
	}, nil)
}
