package usecase

import (
	c "clean/app/domain/common"
	"clean/app/domain/fuga"
	"context"
	"time"
)

func (u *fugaUsecase) Update(ctx context.Context, input fuga.UpdateUsecaseInput) (*fuga.UpdateUsecaseOutput, *c.Error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, repoErr := u.fugaRepository.Find(timeOutCtx, input.Id)
	if repoErr != nil {
		return nil, c.NewInternalServerError(repoErr, "")
	}

	props := fuga.UpdateProps{
		Name:   input.Name,
		Number: input.Number,
	}
	result.Update(props)

	result, repoErr = u.fugaRepository.Update(timeOutCtx, result)
	if repoErr != nil {
		return nil, c.NewInternalServerError(repoErr, "")
	}

	output := new(fuga.UpdateUsecaseOutput)
	output.Fuga = result
	return output, nil
}
