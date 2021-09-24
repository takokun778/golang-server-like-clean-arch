package usecase

import (
	c "clean/app/domain/common"
	"clean/app/domain/hoge"
	"context"
	"time"
)

func (u *hogeUsecase) Create(ctx context.Context, input hoge.CreateUsecaseInput) (*hoge.CreateUsecaseOutput, *c.Error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	props := hoge.CreateNewProps(input)
	result := hoge.CreateNew(props)

	result, repoErr := u.hogeRepository.Save(timeOutCtx, result)
	if repoErr != nil {
		return nil, c.NewInternalServerError(repoErr, "")
	}

	output := new(hoge.CreateUsecaseOutput)
	output.Hoge = result
	return output, nil
}
