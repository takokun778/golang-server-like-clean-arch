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

func TestFugaUsecaseCreate(t *testing.T) {
	var ctx = context.Background()
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockResult := fuga.Create("fuga", 1)

		mmg := mf.NewMockRepository(ctrl)

		mmg.EXPECT().Save(gomock.Any(), gomock.Any()).Return(mockResult, nil)

		usecase := fu.NewFugaUsecase(mmg)

		input := fuga.UsecaseCreateInput{
			Name:   "fuga",
			Number: 1,
		}

		result, err := usecase.Create(ctx, input)

		if err != nil {
			log.Fatal(err)
		}

		assert.Equal(t, result.Fuga.Name(), mockResult.Name())
		assert.Equal(t, result.Fuga.Number(), mockResult.Number())
	})
}
