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

func TestFugaUsecaseFetchAll(t *testing.T) {
	var ctx = context.Background()
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		list := make([]*fuga.Fuga, 0)

		for i := 0; i < 5; i++ {
			list = append(list, fuga.CreateNew("fuga", 1))
		}

		mockResult := fuga.NewList(list)

		mmg := mf.NewMockRepository(ctrl)

		mmg.EXPECT().FindAll(gomock.Any()).Return(mockResult, nil)

		usecase := fu.NewFugaUsecase(mmg)

		input := fuga.UsecaseFetchAllInput{}

		result, err := usecase.FetchAll(ctx, input)

		if err != nil {
			log.Fatal(err)
		}

		assert.Equal(t, result.FugaList, mockResult)
	})
}
