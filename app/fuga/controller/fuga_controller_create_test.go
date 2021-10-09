package controller_test

import (
	"clean/app/domain/fuga"
	fc "clean/app/fuga/controller"
	mf "clean/mock/fuga"
	pbFuga "clean/proto/fuga"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFugaControllerCreate(t *testing.T) {
	var ctx = context.Background()
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		model := fuga.CreateNew("fuga", 1)

		mfu := mf.NewMockUsecase(ctrl)

		mockInput := fuga.UsecaseCreateInput{
			Name:   "fuga",
			Number: 1,
		}

		mockResult := fuga.UsecaseCreateOutput{
			Fuga: model,
		}

		mfu.EXPECT().Create(gomock.Any(), mockInput).Return(&mockResult, nil)

		controller := fc.NewFugaController(mfu)

		req := new(pbFuga.CreateRequest)
		req.Name = "fuga"
		req.Number = 1

		res, err := controller.Create(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, res.Fuga.Name, req.Name)
		assert.Equal(t, res.Fuga.Number, req.Number)
	})
}
