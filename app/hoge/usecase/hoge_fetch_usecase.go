package usecase

import (
	c "clean/app/domain/common"
	"clean/app/domain/hoge"
	"context"
	"time"
)

func (u *hogeUsecase) Fetch(ctx context.Context, input hoge.FetchUsecaseInput) (*hoge.FetchUsecaseOutput, *c.Error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, repoErr := u.hogeRepository.Find(timeOutCtx, input.Id)
	if repoErr != nil {
		return nil, c.NewInternalServerError(repoErr, "")
	}

	output := new(hoge.FetchUsecaseOutput)
	output.Hoge = result
	return output, nil
}
