package usecase

import (
	c "clean/app/domain/common"
	"clean/app/domain/fuga"
	"context"
	"time"
)

func (u *fugaUsecase) FetchAll(ctx context.Context, input fuga.FetchAllUsecaseInput) (*fuga.FetchAllUsecaseOutput, *c.Error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, repoErr := u.fugaRepository.FindAll(timeOutCtx)
	if repoErr != nil {
		return nil, c.NewInternalServerError(repoErr, "")
	}

	output := new(fuga.FetchAllUsecaseOutput)
	output.FugaList = result
	return output, nil
}
