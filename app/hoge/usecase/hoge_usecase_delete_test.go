package usecase_test

import (
	"context"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"clean/app/domain/hoge"
	hu "clean/app/hoge/usecase"
	mh "clean/mock/hoge"
)

func TestHogeUsecaseDelete(t *testing.T) {
	var ctx = context.Background()
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockResult := hoge.Create("hoge", 1).Values()

		mmg := mh.NewMockRepository(ctrl)

		mmg.EXPECT().Find(gomock.Any(), mockResult.Id()).Return(mockResult, nil)
		mmg.EXPECT().Delete(gomock.Any(), mockResult.Id()).Return(nil)

		usecase := hu.NewHogeUsecase(mmg)

		input := hoge.UsecaseDeleteInput{
			Id: mockResult.Id(),
		}

		result, err := usecase.Delete(ctx, input)

		if err != nil {
			log.Fatal(err)
		}

		assert.Equal(t, result.Hoge, mockResult)
	})
}
