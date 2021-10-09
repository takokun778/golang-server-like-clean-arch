package controller_test

import (
	"clean/app/domain/hoge"
	hc "clean/app/hoge/controller"
	mh "clean/mock/hoge"
	pbHoge "clean/proto/hoge"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHogeControllerCreate(t *testing.T) {
	var ctx = context.Background()
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		model := hoge.CreateNew("hoge", 1)

		mhu := mh.NewMockUsecase(ctrl)

		mockInput := hoge.UsecaseCreateInput{
			Name:   "hoge",
			Number: 1,
		}

		mockResult := hoge.UsecaseCreateOutput{
			Hoge: model,
		}

		mhu.EXPECT().Create(gomock.Any(), mockInput).Return(&mockResult, nil)

		controller := hc.NewHogeController(mhu)

		req := new(pbHoge.CreateRequest)
		req.Name = "hoge"
		req.Number = 1

		res, err := controller.Create(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, res.Hoge.Name, req.Name)
		assert.Equal(t, res.Hoge.Number, req.Number)
	})
}
