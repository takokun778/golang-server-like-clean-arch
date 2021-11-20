package usecase

import (
	"clean/app/domain/common"
	"clean/app/domain/fuga"
	"context"
	"time"
)

func (u *fugaUsecase) Delete(ctx context.Context, input fuga.UsecaseDeleteInput) (fuga.UsecaseDeleteOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.fugaRepository.Find(timeOutCtx, input.Id)

	if err != nil {
		return fuga.UsecaseDeleteOutput{}, common.NewInternalServerError(err, "")
	}

	err = u.fugaRepository.Delete(timeOutCtx, input.Id)

	if err != nil {
		return fuga.UsecaseDeleteOutput{}, common.NewInternalServerError(err, "")
	}

	return fuga.UsecaseDeleteOutput{
		Fuga: result,
	}, nil
}
