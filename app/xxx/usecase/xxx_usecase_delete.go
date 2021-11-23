package usecase

import (
	"context"
	"time"

	xe "xxx/app/domain/error"
	"xxx/app/domain/xxx"
)

func (u *xxxUsecase) Delete(ctx context.Context, input *xxx.UsecaseDeleteInput) (*xxx.UsecaseDeleteOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.xxxRepository.Find(timeOutCtx, &xxx.RepositoryFindInput{Id: input.Id})

	if err != nil {
		return nil, xe.NewInternalServerError(err, "")
	}

	_, err = u.xxxRepository.Delete(timeOutCtx, &xxx.RepositoryDeleteInput{Id: input.Id})

	if err != nil {
		return nil, xe.NewInternalServerError(err, "")
	}

	return &xxx.UsecaseDeleteOutput{
		Xxx: result.Xxx,
	}, nil
}
