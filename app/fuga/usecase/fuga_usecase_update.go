package usecase

import (
	"clean/app/domain/common"
	"clean/app/domain/fuga"
	"context"
	"time"
)

func (u *fugaUsecase) Update(ctx context.Context, input fuga.UsecaseUpdateInput) (fuga.UsecaseUpdateOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	src, err := u.fugaRepository.Find(timeOutCtx, input.Id)

	if err != nil {
		return fuga.UsecaseUpdateOutput{}, common.NewInternalServerError(err, "")
	}

	dst := fuga.Reconstruct(src).Update(input.Name, input.Number)

	_, err = u.fugaRepository.Update(timeOutCtx, dst.Values())

	if err != nil {
		return fuga.UsecaseUpdateOutput{}, common.NewInternalServerError(err, "")
	}

	return fuga.UsecaseUpdateOutput{
		Fuga: dst.Values(),
	}, nil
}
