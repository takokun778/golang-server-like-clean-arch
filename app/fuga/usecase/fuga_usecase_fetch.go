package usecase

import (
	c "clean/app/domain/common"
	"clean/app/domain/fuga"
	"context"
	"time"
)

func (u *fugaUsecase) Fetch(ctx context.Context, input fuga.UsecaseFetchInput) (*fuga.UsecaseFetchOutput, *c.Error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, repoErr := u.fugaRepository.Find(timeOutCtx, input.Id)
	if repoErr != nil {
		return nil, c.NewInternalServerError(repoErr, "")
	}

	output := new(fuga.UsecaseFetchOutput)
	output.Fuga = result
	return output, nil
}
