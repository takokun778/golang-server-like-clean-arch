package usecase

import (
	"context"
	"time"

	xe "xxx/app/domain/error"
	"xxx/app/domain/xxx"
)

func (u *xxxUsecase) Update(ctx context.Context, input *xxx.UsecaseUpdateInput) (*xxx.UsecaseUpdateOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	src, err := u.xxxRepository.Find(timeOutCtx, &xxx.RepositoryFindInput{Id: input.Id})

	if err != nil {
		return nil, xe.NewInternalServerError(err, "")
	}

	dst := xxx.Reconstruct(src.Xxx).Update(input.Name, input.Number)

	_, err = u.xxxRepository.Update(timeOutCtx, &xxx.RepositoryUpdateInput{Xxx: dst.Props()})

	if err != nil {
		return nil, xe.NewInternalServerError(err, "")
	}

	return &xxx.UsecaseUpdateOutput{
		Xxx: dst.Props(),
	}, nil
}
