package usecase

import (
	"clean/app/domain/common"
	"clean/app/domain/fuga"
	"context"
	"time"
)

func (u *fugaUsecase) FetchAll(ctx context.Context, input fuga.UsecaseFetchAllInput) (fuga.UsecaseFetchAllOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.fugaRepository.FindAll(timeOutCtx)

	if err != nil {
		return fuga.UsecaseFetchAllOutput{}, common.NewInternalServerError(err, "")
	}

	return fuga.UsecaseFetchAllOutput{
		FugaList: result,
	}, nil
}
