package usecase

import (
	c "clean/app/domain/common"
	"clean/app/domain/fuga"
	"context"
	"time"
)

func (u *fugaUsecase) Update(ctx context.Context, input fuga.UsecaseUpdateInput) (*fuga.UsecaseUpdateOutput, *c.Error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, repoErr := u.fugaRepository.Find(timeOutCtx, input.Id)
	if repoErr != nil {
		return nil, c.NewInternalServerError(repoErr, "")
	}

	result.Update(input.Name, input.Number)

	result, repoErr = u.fugaRepository.Update(timeOutCtx, result)
	if repoErr != nil {
		return nil, c.NewInternalServerError(repoErr, "")
	}

	output := new(fuga.UsecaseUpdateOutput)
	output.Fuga = result
	return output, nil
}
