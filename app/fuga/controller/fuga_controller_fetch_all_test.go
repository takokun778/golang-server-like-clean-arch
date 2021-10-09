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

func TestFugaControllerFetchAll(t *testing.T) {
	var ctx = context.Background()
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		list := make([]*fuga.Fuga, 0)

		for i := 0; i < 5; i++ {
			list = append(list, fuga.CreateNew("fuga", 1))
		}

		mfu := mf.NewMockUsecase(ctrl)

		mockInput := fuga.UsecaseFetchAllInput{}

		mockResult := fuga.UsecaseFetchAllOutput{
			FugaList: fuga.NewList(list),
		}

		mfu.EXPECT().FetchAll(gomock.Any(), mockInput).Return(&mockResult, nil)

		controller := fc.NewFugaController(mfu)

		req := new(pbFuga.FetchAllRequest)

		res, err := controller.FetchAll(ctx, req)

		assert.NoError(t, err)
		assert.Len(t, res.Fugas, 5)
	})
}
