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

func TestHogeUsecaseFetchAll(t *testing.T) {
	var ctx = context.Background()
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		list := make([]*hoge.Hoge, 0)

		for i := 0; i < 5; i++ {
			list = append(list, hoge.CreateNew("hoge", 1))
		}

		mockResult := hoge.NewList(list)

		mmg := mh.NewMockRepository(ctrl)

		mmg.EXPECT().FindAll(gomock.Any()).Return(mockResult, nil)

		usecase := hu.NewHogeUsecase(mmg)

		input := hoge.UsecaseFetchAllInput{}

		result, err := usecase.FetchAll(ctx, input)

		if err != nil {
			log.Fatal(err)
		}

		assert.Equal(t, result.HogeList, mockResult)
	})
}
