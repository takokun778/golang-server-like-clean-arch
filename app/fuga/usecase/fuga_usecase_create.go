package usecase

import (
	"clean/app/domain/common"
	"clean/app/domain/fuga"
	"context"
	"time"
)

func (u *fugaUsecase) Create(ctx context.Context, input fuga.UsecaseCreateInput) (*fuga.UsecaseCreateOutput, *common.Error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result := fuga.CreateNew(input.Name, input.Number)

	result, repoErr := u.fugaRepository.Save(timeOutCtx, result)
	if repoErr != nil {
		return nil, common.NewInternalServerError(repoErr, "")
	}

	output := new(fuga.UsecaseCreateOutput)
	output.Fuga = result
	return output, nil
}
