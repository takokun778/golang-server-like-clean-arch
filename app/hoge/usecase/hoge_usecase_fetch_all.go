package usecase

import (
	"clean/app/domain/common"
	"clean/app/domain/hoge"
	"context"
	"time"
)

func (u *hogeUsecase) FetchAll(ctx context.Context, input hoge.UsecaseFetchAllInput) (*hoge.UsecaseFetchAllOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.hogeRepository.FindAll(timeOutCtx)
	if err != nil {
		return nil, common.NewInternalServerError(err, "")
	}

	output := new(hoge.UsecaseFetchAllOutput)
	output.HogeList = result
	return output, nil
}
