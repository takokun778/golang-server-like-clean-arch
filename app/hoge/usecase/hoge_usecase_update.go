package usecase

import (
	c "clean/app/domain/common"
	"clean/app/domain/hoge"
	"context"
	"time"
)

func (u *hogeUsecase) Update(ctx context.Context, input hoge.UsecaseUpdateInput) (*hoge.UsecaseUpdateOutput, *c.Error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, repoErr := u.hogeRepository.Find(timeOutCtx, input.Id)
	if repoErr != nil {
		return nil, c.NewInternalServerError(repoErr, "")
	}

	result.Update(input.Name, input.Number)

	result, repoErr = u.hogeRepository.Update(timeOutCtx, result)
	if repoErr != nil {
		return nil, c.NewInternalServerError(repoErr, "")
	}

	output := new(hoge.UsecaseUpdateOutput)
	output.Hoge = result
	return output, nil
}