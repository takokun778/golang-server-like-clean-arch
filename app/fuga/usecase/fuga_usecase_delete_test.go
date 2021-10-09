package usecase_test

import (
	"context"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"clean/app/domain/fuga"
	fu "clean/app/fuga/usecase"
	mf "clean/mock/fuga"
)

func TestFugaUsecaseDelete(t *testing.T) {
	var ctx = context.Background()
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockResult := fuga.CreateNew("fuga", 1)

		mmg := mf.NewMockRepository(ctrl)

		mmg.EXPECT().Find(gomock.Any(), mockResult.Id()).Return(mockResult, nil)
		mmg.EXPECT().Delete(gomock.Any(), mockResult.Id()).Return(nil)

		usecase := fu.NewFugaUsecase(mmg)

		input := fuga.UsecaseDeleteInput{
			Id: mockResult.Id(),
		}

		result, err := usecase.Delete(ctx, input)

		if err != nil {
			log.Fatal(err)
		}

		assert.Equal(t, result.Fuga, mockResult)
	})
}
