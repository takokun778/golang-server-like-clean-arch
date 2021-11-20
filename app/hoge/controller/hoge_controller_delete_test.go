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

func TestHogeControllerDelete(t *testing.T) {
	var ctx = context.Background()
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		model := hoge.Create("hoge", 1)

		mhu := mh.NewMockUsecase(ctrl)

		mockInput := hoge.UsecaseDeleteInput{
			Id: model.Id(),
		}

		mockResult := hoge.UsecaseDeleteOutput{
			Hoge: model.Values(),
		}

		mhu.EXPECT().Delete(gomock.Any(), mockInput).Return(&mockResult, nil)

		controller := hc.NewHogeController(mhu)

		req := new(pbHoge.DeleteRequest)
		req.Id = model.Id().Value().String()

		_, err := controller.Delete(ctx, req)

		assert.NoError(t, err)
	})
}
