package interactor

import (
	"context"
	"time"

	dErr "xxx/app/domain/error"
	"xxx/app/domain/xxx"
)

func (i *xxxInteractor) ReadAll(ctx context.Context, input *xxx.UsecaseReadAllInput) (*xxx.UsecaseReadAllOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := i.xxxRepository.FindAll(timeOutCtx, &xxx.RepositoryFindAllItem{})

	if err != nil {
		return nil, dErr.NewInternalServerError(err, "")
	}

	return &xxx.UsecaseReadAllOutput{
		Xxxs: result,
	}, nil
}
