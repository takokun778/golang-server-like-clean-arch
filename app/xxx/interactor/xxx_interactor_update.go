package interactor

import (
	"context"
	"time"

	dError "xxx/app/domain/error"
	"xxx/app/domain/xxx"
)

func (i *xxxInteractor) Update(ctx context.Context, input *xxx.UsecaseUpdateInput) (*xxx.UsecaseUpdateOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	src, err := i.xxxRepository.Find(timeOutCtx, &xxx.RepositoryFindItem{Id: input.Id})

	if err != nil {
		return nil, dError.NewInternalServerError(err, "")
	}

	dst := xxx.Reconstruct(src).Update(input.Name, input.Number)

	_, err = i.xxxRepository.Update(timeOutCtx, &xxx.RepositoryUpdateItem{Xxx: dst.Props()})

	if err != nil {
		return nil, dError.NewInternalServerError(err, "")
	}

	return &xxx.UsecaseUpdateOutput{
		Xxx: dst.Props(),
	}, nil
}
