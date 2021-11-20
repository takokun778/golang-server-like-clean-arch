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

func TestHogeUsecaseCreate(t *testing.T) {
	var ctx = context.Background()
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockResult := hoge.Create("hoge", 1)

		mmg := mh.NewMockRepository(ctrl)

		mmg.EXPECT().Save(gomock.Any(), gomock.Any()).Return(mockResult, nil)

		usecase := hu.NewHogeUsecase(mmg)

		input := hoge.UsecaseCreateInput{
			Name:   "hoge",
			Number: 1,
		}

		result, err := usecase.Create(ctx, input)

		if err != nil {
			log.Fatal(err)
		}

		assert.Equal(t, result.Hoge.Name(), mockResult.Name())
		assert.Equal(t, result.Hoge.Number(), mockResult.Number())
	})
}
