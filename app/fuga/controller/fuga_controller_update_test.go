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

func TestFugaControllerUpdate(t *testing.T) {
	var ctx = context.Background()
	t.Run("normal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		model := fuga.CreateNew("fuga", 1)
		model.Update("fugafuga", 2)

		mfu := mf.NewMockUsecase(ctrl)

		mockInput := fuga.UsecaseUpdateInput{
			Id:     model.Id(),
			Name:   "fugafuga",
			Number: 2,
		}

		mockResult := fuga.UsecaseUpdateOutput{
			Fuga: model,
		}

		mfu.EXPECT().Update(gomock.Any(), mockInput).Return(&mockResult, nil)

		controller := fc.NewFugaController(mfu)

		req := new(pbFuga.UpdateRequest)
		req.Id = model.Id().Value().String()
		req.Name = "fugafuga"
		req.Number = 2

		res, err := controller.Update(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, res.Fuga.Id, req.Id)
		assert.Equal(t, res.Fuga.Name, req.Name)
		assert.Equal(t, res.Fuga.Number, req.Number)
	})
}
