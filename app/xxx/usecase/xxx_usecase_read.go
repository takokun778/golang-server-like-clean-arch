package usecase

import (
	"context"
	"time"

	xe "xxx/app/domain/error"
	"xxx/app/domain/xxx"
)

func (u *xxxUsecase) Read(ctx context.Context, input *xxx.UsecaseReadInput) (*xxx.UsecaseReadOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.xxxRepository.Find(timeOutCtx, &xxx.RepositoryFindInput{Id: input.Id})

	if err != nil {
		return nil, xe.NewInternalServerError(err, "")
	}

	return &xxx.UsecaseReadOutput{
		Xxx: result.Xxx,
	}, nil
}
