package usecase

import (
	"clean/app/domain/common"
	"clean/app/domain/fuga"
	"context"
	"time"
)

func (u *fugaUsecase) Update(ctx context.Context, input fuga.UsecaseUpdateInput) (*fuga.UsecaseUpdateOutput, *common.Error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	src, err := u.fugaRepository.Find(timeOutCtx, input.Id)
	if err != nil {
		return nil, common.NewInternalServerError(err, "")
	}

	dst := src.Update(input.Name, input.Number)

	dst, err = u.fugaRepository.Update(timeOutCtx, dst)
	if err != nil {
		return nil, common.NewInternalServerError(err, "")
	}

	output := new(fuga.UsecaseUpdateOutput)
	output.Fuga = dst
	return output, nil
}
