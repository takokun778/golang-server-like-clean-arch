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

func TestHogeControllerFetchAll(t *testing.T) {
	var ctx = context.Background()
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		list := make([]*hoge.Hoge, 0)

		for i := 0; i < 5; i++ {
			list = append(list, hoge.CreateNew("hoge", 1))
		}

		mhu := mh.NewMockUsecase(ctrl)

		mockInput := hoge.UsecaseFetchAllInput{}

		mockResult := hoge.UsecaseFetchAllOutput{
			HogeList: hoge.NewList(list),
		}

		mhu.EXPECT().FetchAll(gomock.Any(), mockInput).Return(&mockResult, nil)

		controller := hc.NewHogeController(mhu)

		req := new(pbHoge.FetchAllRequest)

		res, err := controller.FetchAll(ctx, req)

		assert.NoError(t, err)
		assert.Len(t, res.Hoges, 5)
	})
}
