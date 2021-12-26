package interactor

import (
	"context"
	"time"

	dError "xxx/app/domain/error"
	"xxx/app/domain/xxx"
	"xxx/app/xxx/presenter"
)

func (i *xxxInteractor) Update(ctx context.Context, dto *xxx.UsecaseUpdateDto) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	src, err := i.xxxRepository.Find(timeOutCtx, &xxx.RepositoryFindItem{Id: dto.Id})

	if err != nil {
		i.xxxPresenter.Update(ctx, nil, dError.NewInternalServerError(err, ""))
		return
	}

	dst := xxx.Reconstruct(src).Update(dto.Name, dto.Number)

	_, err = i.xxxRepository.Update(timeOutCtx, &xxx.RepositoryUpdateItem{Xxx: dst.Props()})

	if err != nil {
		i.xxxPresenter.Update(ctx, nil, dError.NewInternalServerError(err, ""))
		return
	}

	i.xxxPresenter.Update(ctx, &presenter.PresenterUpdateDto{
		Xxx: dst.Props(),
	}, nil)
}
