package usecase

import (
	c "clean/app/domain/common"
	"clean/app/domain/hoge"
	"context"
	"time"
)

func (u *hogeUsecase) Update(ctx context.Context, input hoge.UpdateUsecaseInput) (*hoge.UpdateUsecaseOutput, *c.Error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, repoErr := u.hogeRepository.Find(timeOutCtx, input.Id)
	if repoErr != nil {
		return nil, c.NewInternalServerError(repoErr, "")
	}

	props := hoge.UpdateProps{
		Name:   input.Name,
		Number: input.Number,
	}
	result.Update(props)

	result, repoErr = u.hogeRepository.Update(timeOutCtx, result)
	if repoErr != nil {
		return nil, c.NewInternalServerError(repoErr, "")
	}

	output := new(hoge.UpdateUsecaseOutput)
	output.Hoge = result
	return output, nil
}
