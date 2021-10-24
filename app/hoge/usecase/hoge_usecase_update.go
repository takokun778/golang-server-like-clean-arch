package usecase

import (
	"clean/app/domain/common"
	"clean/app/domain/hoge"
	"context"
	"time"
)

func (u *hogeUsecase) Update(ctx context.Context, input hoge.UsecaseUpdateInput) (*hoge.UsecaseUpdateOutput, *common.Error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	src, err := u.hogeRepository.Find(timeOutCtx, input.Id)
	if err != nil {
		return nil, common.NewInternalServerError(err, "")
	}

	dst := src.Update(input.Name, input.Number)

	dst, err = u.hogeRepository.Update(timeOutCtx, dst)
	if err != nil {
		return nil, common.NewInternalServerError(err, "")
	}

	output := new(hoge.UsecaseUpdateOutput)
	output.Hoge = dst
	return output, nil
}
