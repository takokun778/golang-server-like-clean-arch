package usecase

import (
	"clean/app/domain/common"
	"clean/app/domain/hoge"
	"context"
	"time"
)

func (u *hogeUsecase) Delete(ctx context.Context, input hoge.UsecaseDeleteInput) (*hoge.UsecaseDeleteOutput, *common.Error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.hogeRepository.Find(timeOutCtx, input.Id)
	if err != nil {
		return nil, common.NewInternalServerError(err, "")
	}

	err = u.hogeRepository.Delete(timeOutCtx, input.Id)
	if err != nil {
		return nil, common.NewInternalServerError(err, "")
	}

	output := new(hoge.UsecaseDeleteOutput)
	output.Hoge = result
	return output, nil
}
