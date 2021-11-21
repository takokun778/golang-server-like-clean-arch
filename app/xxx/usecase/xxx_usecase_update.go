package usecase

import (
	"context"
	"time"

	"xxx/app/domain/common"
	"xxx/app/domain/xxx"
)

func (u *xxxUsecase) Update(ctx context.Context, input *xxx.UsecaseUpdateInput) (*xxx.UsecaseUpdateOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	src, err := u.xxxRepository.Find(timeOutCtx, input.Id)

	if err != nil {
		return nil, common.NewInternalServerError(err, "")
	}

	dst := xxx.Reconstruct(src).Update(input.Name, input.Number)

	_, err = u.xxxRepository.Update(timeOutCtx, dst.Props())

	if err != nil {
		return nil, common.NewInternalServerError(err, "")
	}

	return &xxx.UsecaseUpdateOutput{
		Xxx: dst.Props(),
	}, nil
}
