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

func TestHogeUsecaseUpdate(t *testing.T) {
	var ctx = context.Background()
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockResult := hoge.CreateNew("hoge", 1)
		mockFindResult := hoge.New(mockResult.Id(), mockResult.Name(), mockResult.Number(), mockResult.CreatedAt(), mockResult.UpdatedAt())
		mockResult.Update("hogehoge", 2)

		mmg := mh.NewMockRepository(ctrl)

		mmg.EXPECT().Find(gomock.Any(), mockFindResult.Id()).Return(mockFindResult, nil)
		mmg.EXPECT().Update(gomock.Any(), gomock.Any()).Return(mockResult, nil)

		usecase := hu.NewHogeUsecase(mmg)

		input := hoge.UsecaseUpdateInput{
			Id:     mockFindResult.Id(),
			Name:   hoge.Name("hogehoge"),
			Number: 2,
		}

		result, err := usecase.Update(ctx, input)

		if err != nil {
			log.Fatal(err)
		}

		assert.Equal(t, result.Hoge.Id(), mockResult.Id())
		assert.Equal(t, result.Hoge.Name(), mockResult.Name())
		assert.Equal(t, result.Hoge.Number(), mockResult.Number())
	})
}
