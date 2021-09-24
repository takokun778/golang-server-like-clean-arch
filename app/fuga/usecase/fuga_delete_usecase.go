package usecase

import (
	c "clean/app/domain/common"
	"clean/app/domain/fuga"
	"context"
	"time"
)

func (u *fugaUsecase) Delete(ctx context.Context, input fuga.DeleteUsecaseInput) (*fuga.DeleteUsecaseOutput, *c.Error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, repoErr := u.fugaRepository.Find(timeOutCtx, input.Id)
	if repoErr != nil {
		return nil, c.NewInternalServerError(repoErr, "")
	}

	repoErr = u.fugaRepository.Delete(timeOutCtx, input.Id)
	if repoErr != nil {
		return nil, c.NewInternalServerError(repoErr, "")
	}

	output := new(fuga.DeleteUsecaseOutput)
	output.Fuga = result
	return output, nil
}
