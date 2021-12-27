package interactor

import (
	"context"
	"time"

	dErr "xxx/app/domain/error"
	"xxx/app/domain/xxx"
	"xxx/app/xxx/presenter"
)

func (i *xxxInteractor) Delete(ctx context.Context, dto *xxx.UsecaseDeleteDto, err error) {
	if err != nil {
		i.xxxPresenter.Delete(ctx, nil, err)
		return
	}

	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	foundXxx, err := i.xxxRepository.Find(timeOutCtx, &xxx.RepositoryFindItem{Id: dto.Id})

	if err != nil {
		i.xxxPresenter.Delete(ctx, nil, dErr.NewInternalServerError(err, ""))
		return
	}

	err = i.xxxRepository.Delete(timeOutCtx, &xxx.RepositoryDeleteItem{Id: dto.Id})

	if err != nil {
		i.xxxPresenter.Delete(timeOutCtx, nil, dErr.NewInternalServerError(err, ""))
		return
	}

	i.xxxPresenter.Delete(timeOutCtx, &presenter.PresenterDeleteDto{
		Xxx: foundXxx,
	}, nil)
}
