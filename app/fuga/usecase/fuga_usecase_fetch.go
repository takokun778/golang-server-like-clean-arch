package usecase

import (
	"clean/app/domain/common"
	"clean/app/domain/fuga"
	"context"
	"time"
)

func (u *fugaUsecase) Fetch(ctx context.Context, input fuga.UsecaseFetchInput) (*fuga.UsecaseFetchOutput, *common.Error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.fugaRepository.Find(timeOutCtx, input.Id)
	if err != nil {
		return nil, common.NewInternalServerError(err, "")
	}

	output := new(fuga.UsecaseFetchOutput)
	output.Fuga = result
	return output, nil
}
