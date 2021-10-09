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

	src, repoErr := u.hogeRepository.Find(timeOutCtx, input.Id)
	if repoErr != nil {
		return nil, common.NewInternalServerError(repoErr, "")
	}

	dst := src.Update(input.Name, input.Number)

	dst, repoErr = u.hogeRepository.Update(timeOutCtx, dst)
	if repoErr != nil {
		return nil, common.NewInternalServerError(repoErr, "")
	}

	output := new(hoge.UsecaseUpdateOutput)
	output.Hoge = dst
	return output, nil
}
