package usecase

import (
	"clean/app/domain/common"
	"clean/app/domain/fuga"
	"context"
	"time"
)

func (u *fugaUsecase) Create(ctx context.Context, input fuga.UsecaseCreateInput) (*fuga.UsecaseCreateOutput, *common.Error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result := fuga.CreateNew(input.Name, input.Number)

	result, err := u.fugaRepository.Save(timeOutCtx, result)
	if err != nil {
		return nil, common.NewInternalServerError(err, "")
	}

	output := new(fuga.UsecaseCreateOutput)
	output.Fuga = result
	return output, nil
}
