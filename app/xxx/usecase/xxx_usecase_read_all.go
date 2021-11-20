package usecase

import (
	"context"
	"time"

	"xxx/app/domain/common"
	"xxx/app/domain/xxx"
)

func (u *xxxUsecase) ReadAll(ctx context.Context, input *xxx.UsecaseReadAllInput) (*xxx.UsecaseReadAllOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.xxxRepository.FindAll(timeOutCtx)

	if err != nil {
		return nil, common.NewInternalServerError(err, "")
	}

	return &xxx.UsecaseReadAllOutput{
		Xxxs: result,
	}, nil
}
