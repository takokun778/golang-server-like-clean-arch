package interactor

import (
	"context"
	"time"

	dErr "xxx/app/domain/error"
	"xxx/app/domain/xxx"
)

func (i *xxxInteractor) Read(ctx context.Context, input *xxx.UsecaseReadInput) (*xxx.UsecaseReadOutput, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := i.xxxRepository.Find(timeOutCtx, &xxx.RepositoryFindItem{Id: input.Id})

	if err != nil {
		return nil, dErr.NewInternalServerError(err, "")
	}

	return &xxx.UsecaseReadOutput{
		Xxx: result,
	}, nil
}
