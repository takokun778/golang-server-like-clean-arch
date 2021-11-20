package usecase

import (
	"context"
	"time"

	"xxx/app/domain/common"
	"xxx/app/domain/xxx"
)

func (u *xxxUsecase) Delete(ctx context.Context, input *xxx.UsecaseDeleteInput) (*xxx.UsecaseDeleteOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.xxxRepository.Find(timeOutCtx, input.Id)

	if err != nil {
		return nil, common.NewInternalServerError(err, "")
	}

	err = u.xxxRepository.Delete(timeOutCtx, input.Id)

	if err != nil {
		return nil, common.NewInternalServerError(err, "")
	}

	return &xxx.UsecaseDeleteOutput{
		Xxx: result,
	}, nil
}
