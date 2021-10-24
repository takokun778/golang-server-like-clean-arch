package usecase

import (
	"clean/app/domain/common"
	"clean/app/domain/hoge"
	"context"
	"time"
)

func (u *hogeUsecase) Fetch(ctx context.Context, input hoge.UsecaseFetchInput) (*hoge.UsecaseFetchOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.hogeRepository.Find(timeOutCtx, input.Id)
	if err != nil {
		return nil, common.NewInternalServerError(err, "")
	}

	output := new(hoge.UsecaseFetchOutput)
	output.Hoge = result
	return output, nil
}
