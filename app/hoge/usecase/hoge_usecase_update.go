package usecase

import (
	"clean/app/domain/common"
	"clean/app/domain/hoge"
	"context"
	"time"
)

func (u *hogeUsecase) Update(ctx context.Context, input hoge.UsecaseUpdateInput) (hoge.UsecaseUpdateOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	src, err := u.hogeRepository.Find(timeOutCtx, input.Id)

	if err != nil {
		return hoge.UsecaseUpdateOutput{}, common.NewInternalServerError(err, "")
	}

	dst := hoge.Reconstruct(src).Update(input.Name, input.Number)

	_, err = u.hogeRepository.Update(timeOutCtx, dst.Values())

	if err != nil {
		return hoge.UsecaseUpdateOutput{}, common.NewInternalServerError(err, "")
	}

	return hoge.UsecaseUpdateOutput{
		Hoge: dst.Values(),
	}, nil
}
