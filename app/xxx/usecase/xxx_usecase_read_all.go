package usecase

import (
	"context"
	"time"

	xe "xxx/app/domain/error"
	"xxx/app/domain/xxx"
)

func (u *xxxUsecase) ReadAll(ctx context.Context, input *xxx.UsecaseReadAllInput) (*xxx.UsecaseReadAllOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.xxxRepository.FindAll(timeOutCtx, &xxx.RepositoryFindAllInput{})

	if err != nil {
		return nil, xe.NewInternalServerError(err, "")
	}

	return &xxx.UsecaseReadAllOutput{
		Xxxs: result.Xxxs,
	}, nil
}
