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

func TestFugaControllerDelete(t *testing.T) {
	var ctx = context.Background()
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		model := fuga.Create("fuga", 1)

		mfu := mf.NewMockUsecase(ctrl)

		mockInput := fuga.UsecaseDeleteInput{
			Id: model.Id(),
		}

		mockResult := fuga.UsecaseDeleteOutput{
			Fuga: model.Values(),
		}

		mfu.EXPECT().Delete(gomock.Any(), mockInput).Return(&mockResult, nil)

		controller := fc.NewFugaController(mfu)

		req := new(pbFuga.DeleteRequest)
		req.Id = model.Id().Value().String()

		_, err := controller.Delete(ctx, req)

		assert.NoError(t, err)
	})
}
