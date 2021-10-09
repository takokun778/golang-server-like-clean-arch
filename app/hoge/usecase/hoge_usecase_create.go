package usecase

import (
	c "clean/app/domain/common"
	"clean/app/domain/hoge"
	"context"
	"time"
)

func (u *hogeUsecase) Create(ctx context.Context, input hoge.UsecaseCreateInput) (*hoge.UsecaseCreateOutput, *c.Error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result := hoge.CreateNew(input.Name, input.Number)

	result, repoErr := u.hogeRepository.Save(timeOutCtx, result)
	if repoErr != nil {
		return nil, c.NewInternalServerError(repoErr, "")
	}

	output := new(hoge.UsecaseCreateOutput)
	output.Hoge = result
	return output, nil
}