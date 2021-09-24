package usecase

import (
	c "clean/app/domain/common"
	"clean/app/domain/fuga"
	"context"
	"time"
)

func (u *fugaUsecase) Create(ctx context.Context, input fuga.CreateUsecaseInput) (*fuga.CreateUsecaseOutput, *c.Error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result := fuga.CreateNew(input.Name, input.Number)

	result, repoErr := u.fugaRepository.Save(timeOutCtx, result)
	if repoErr != nil {
		return nil, c.NewInternalServerError(repoErr, "")
	}

	output := new(fuga.CreateUsecaseOutput)
	output.Fuga = result
	return output, nil
}
