package interactor

import (
	"context"
	"time"

	dErr "xxx/app/domain/error"
	"xxx/app/domain/xxx"
)

func (i *xxxInteractor) Delete(ctx context.Context, input *xxx.UsecaseDeleteInput) (*xxx.UsecaseDeleteOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := i.xxxRepository.Find(timeOutCtx, &xxx.RepositoryFindItem{Id: input.Id})

	if err != nil {
		return nil, dErr.NewInternalServerError(err, "")
	}

	err = i.xxxRepository.Delete(timeOutCtx, &xxx.RepositoryDeleteItem{Id: input.Id})

	if err != nil {
		return nil, dErr.NewInternalServerError(err, "")
	}

	return &xxx.UsecaseDeleteOutput{
		Xxx: result,
	}, nil
}
