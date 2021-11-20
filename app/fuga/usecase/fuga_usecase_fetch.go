package usecase

import (
	"clean/app/domain/common"
	"clean/app/domain/fuga"
	"context"
	"time"
)

func (u *fugaUsecase) Fetch(ctx context.Context, input fuga.UsecaseFetchInput) (fuga.UsecaseFetchOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.fugaRepository.Find(timeOutCtx, input.Id)

	if err != nil {
		return fuga.UsecaseFetchOutput{}, common.NewInternalServerError(err, "")
	}

	return fuga.UsecaseFetchOutput{
		Fuga: result,
	}, nil
}
