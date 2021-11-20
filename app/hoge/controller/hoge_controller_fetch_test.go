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

func TestHogeControllerFetch(t *testing.T) {
	var ctx = context.Background()
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		model := hoge.Create("hoge", 1)

		mhu := mh.NewMockUsecase(ctrl)

		mockInput := hoge.UsecaseFetchInput{
			Id: model.Id(),
		}

		mockResult := hoge.UsecaseFetchOutput{
			Hoge: model.Values(),
		}

		mhu.EXPECT().Fetch(gomock.Any(), mockInput).Return(&mockResult, nil)

		controller := hc.NewHogeController(mhu)

		req := new(pbHoge.FetchRequest)
		req.Id = model.Id().Value().String()

		res, err := controller.Fetch(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, res.Hoge.Id, model.Id().Value().String())
	})
}
