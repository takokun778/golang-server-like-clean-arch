package usecase

import (
	c "clean/app/domain/common"
	"clean/app/domain/hoge"
	"context"
	"time"
)

func (u *hogeUsecase) Delete(ctx context.Context, input hoge.DeleteUsecaseInput) (*hoge.DeleteUsecaseOutput, *c.Error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, repoErr := u.hogeRepository.Find(timeOutCtx, input.Id)
	if repoErr != nil {
		return nil, c.NewInternalServerError(repoErr, "")
	}

	repoErr = u.hogeRepository.Delete(timeOutCtx, input.Id)
	if repoErr != nil {
		return nil, c.NewInternalServerError(repoErr, "")
	}

	output := new(hoge.DeleteUsecaseOutput)
	output.Hoge = result
	return output, nil
}
