package usecase

import (
	"clean/app/domain/common"
	"clean/app/domain/hoge"
	"context"
	"time"
)

func (u *hogeUsecase) Create(ctx context.Context, input hoge.UsecaseCreateInput) (*hoge.UsecaseCreateOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result := hoge.CreateNew(input.Name, input.Number)

	result, err := u.hogeRepository.Save(timeOutCtx, result)
	if err != nil {
		return nil, common.NewInternalServerError(err, "")
	}

	output := new(hoge.UsecaseCreateOutput)
	output.Hoge = result
	return output, nil
}
