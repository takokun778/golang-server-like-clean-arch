package usecase

import (
	"context"
	"time"

	"xxx/app/domain/common"
	"xxx/app/domain/xxx"
)

func (u *xxxUsecase) Read(ctx context.Context, input *xxx.UsecaseReadInput) (*xxx.UsecaseReadOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.xxxRepository.Find(timeOutCtx, input.Id)

	if err != nil {
		return nil, common.NewInternalServerError(err, "")
	}

	return &xxx.UsecaseReadOutput{
		Xxx: result,
	}, nil
}
