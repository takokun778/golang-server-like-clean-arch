package usecase

import (
	"clean/app/domain/common"
	"clean/app/domain/fuga"
	"context"
	"time"
)

func (u *fugaUsecase) Create(ctx context.Context, input fuga.UsecaseCreateInput) (fuga.UsecaseCreateOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result := fuga.Create(input.Name, input.Number)

	_, err := u.fugaRepository.Save(timeOutCtx, result.Values())

	if err != nil {
		return fuga.UsecaseCreateOutput{}, common.NewInternalServerError(err, "")
	}

	return fuga.UsecaseCreateOutput{
		Fuga: result.Values(),
	}, nil
}
