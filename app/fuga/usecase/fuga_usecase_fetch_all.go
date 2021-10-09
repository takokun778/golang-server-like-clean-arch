package usecase

import (
	"clean/app/domain/common"
	"clean/app/domain/fuga"
	"context"
	"time"
)

func (u *fugaUsecase) FetchAll(ctx context.Context, input fuga.UsecaseFetchAllInput) (*fuga.UsecaseFetchAllOutput, *common.Error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, repoErr := u.fugaRepository.FindAll(timeOutCtx)
	if repoErr != nil {
		return nil, common.NewInternalServerError(repoErr, "")
	}

	output := new(fuga.UsecaseFetchAllOutput)
	output.FugaList = result
	return output, nil
}
