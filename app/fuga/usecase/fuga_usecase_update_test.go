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

func TestFugaUsecaseUpdate(t *testing.T) {
	var ctx = context.Background()
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockResult := fuga.CreateNew("fuga", 1)
		mockFindResult := fuga.New(mockResult.Id(), mockResult.Name(), mockResult.Number(), mockResult.CreatedAt(), mockResult.UpdatedAt())
		mockResult.Update("fugafuga", 2)

		mmg := mf.NewMockRepository(ctrl)

		mmg.EXPECT().Find(gomock.Any(), mockFindResult.Id()).Return(mockFindResult, nil)
		mmg.EXPECT().Update(gomock.Any(), gomock.Any()).Return(mockResult, nil)

		usecase := fu.NewFugaUsecase(mmg)

		input := fuga.UsecaseUpdateInput{
			Id:     mockFindResult.Id(),
			Name:   fuga.Name("fugafuga"),
			Number: 2,
		}

		result, err := usecase.Update(ctx, input)

		if err != nil {
			log.Fatal(err)
		}

		assert.Equal(t, result.Fuga.Id(), mockResult.Id())
		assert.Equal(t, result.Fuga.Name(), mockResult.Name())
		assert.Equal(t, result.Fuga.Number(), mockResult.Number())
	})
}
