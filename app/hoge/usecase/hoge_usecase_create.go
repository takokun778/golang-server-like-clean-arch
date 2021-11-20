package usecase

import (
	"clean/app/domain/common"
	"clean/app/domain/hoge"
	"context"
	"time"
)

func (u *hogeUsecase) Create(ctx context.Context, input hoge.UsecaseCreateInput) (hoge.UsecaseCreateOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result := hoge.Create(input.Name, input.Number)

	_, err := u.hogeRepository.Save(timeOutCtx, result.Values())
	if err != nil {
		return hoge.UsecaseCreateOutput{}, common.NewInternalServerError(err, "")
	}

	return hoge.UsecaseCreateOutput{
		Hoge: result.Values(),
	}, nil
}
