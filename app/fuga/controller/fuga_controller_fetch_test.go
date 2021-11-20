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

func TestFugaControllerFetch(t *testing.T) {
	var ctx = context.Background()
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		model := fuga.Create("fuga", 1)

		mfu := mf.NewMockUsecase(ctrl)

		mockInput := fuga.UsecaseFetchInput{
			Id: model.Id(),
		}

		mockResult := fuga.UsecaseFetchOutput{
			Fuga: model.Values(),
		}

		mfu.EXPECT().Fetch(gomock.Any(), mockInput).Return(&mockResult, nil)

		controller := fc.NewFugaController(mfu)

		req := new(pbFuga.FetchRequest)
		req.Id = model.Id().Value().String()

		res, err := controller.Fetch(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, res.Fuga.Id, model.Id().Value().String())
	})
}
