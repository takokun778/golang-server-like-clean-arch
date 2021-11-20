package usecase

import (
	"context"
	"time"

	"xxx/app/domain/common"
	"xxx/app/domain/xxx"
)

func (u *xxxUsecase) Create(ctx context.Context, input *xxx.UsecaseCreateInput) (*xxx.UsecaseCreateOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result := xxx.Create(input.Name, input.Number)

	_, err := u.xxxRepository.Save(timeOutCtx, result.Values())

	if err != nil {
		return nil, common.NewInternalServerError(err, "")
	}

	return &xxx.UsecaseCreateOutput{
		Xxx: result.Values(),
	}, nil
}
