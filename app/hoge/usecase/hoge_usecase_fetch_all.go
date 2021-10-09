package usecase

import (
	c "clean/app/domain/common"
	"clean/app/domain/hoge"
	"context"
	"time"
)

func (u *hogeUsecase) FetchAll(ctx context.Context, input hoge.UsecaseFetchAllInput) (*hoge.UsecaseFetchAllOutput, *c.Error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, repoErr := u.hogeRepository.FindAll(timeOutCtx)
	if repoErr != nil {
		return nil, c.NewInternalServerError(repoErr, "")
	}

	output := new(hoge.UsecaseFetchAllOutput)
	output.HogeList = result
	return output, nil
}
