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

func TestHogeControllerUpdate(t *testing.T) {
	var ctx = context.Background()
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		model := hoge.CreateNew("hoge", 1)
		model.Update("hogehoge", 2)

		mhu := mh.NewMockUsecase(ctrl)

		mockInput := hoge.UsecaseUpdateInput{
			Id:     model.Id(),
			Name:   "hogehoge",
			Number: 2,
		}

		mockResult := hoge.UsecaseUpdateOutput{
			Hoge: model,
		}

		mhu.EXPECT().Update(gomock.Any(), mockInput).Return(&mockResult, nil)

		controller := hc.NewHogeController(mhu)

		req := new(pbHoge.UpdateRequest)
		req.Id = model.Id().Value().String()
		req.Name = "hogehoge"
		req.Number = 2

		res, err := controller.Update(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, res.Hoge.Id, req.Id)
		assert.Equal(t, res.Hoge.Name, req.Name)
		assert.Equal(t, res.Hoge.Number, req.Number)
	})
}
