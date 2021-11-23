package usecase

import (
	"context"
	"time"

	dErr "xxx/app/domain/error"
	"xxx/app/domain/xxx"
)

func (u *xxxUsecase) Delete(ctx context.Context, input *xxx.UsecaseDeleteInput) (*xxx.UsecaseDeleteOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.xxxRepository.Find(timeOutCtx, &xxx.RepositoryFindItem{Id: input.Id})

	if err != nil {
		return nil, dErr.NewInternalServerError(err, "")
	}

	err = u.xxxRepository.Delete(timeOutCtx, &xxx.RepositoryDeleteItem{Id: input.Id})

	if err != nil {
		return nil, dErr.NewInternalServerError(err, "")
	}

	return &xxx.UsecaseDeleteOutput{
		Xxx: result,
	}, nil
}
