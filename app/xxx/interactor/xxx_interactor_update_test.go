package interactor_test

import (
	"context"
	"testing"

	dErr "xxx/app/domain/error"
	"xxx/app/domain/xxx"
	xi "xxx/app/xxx/interactor"
	mx "xxx/mock/xxx"
	"xxx/test"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestXxxInteractorUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mxg := mx.NewMockRepository(ctrl)

	tests := make([]test.Case, 0)

	test1Xxx := test.NewXxx("test1", 1)
	test1Name, _ := xxx.NewName("updated")
	test1Number, _ := xxx.NewNumber(2)
	test1XxxUpdated := test1Xxx.Update(test1Name, test1Number)
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			mxg.EXPECT().Find(gomock.Any(), &xxx.RepositoryFindItem{Id: test1Xxx.Id()}).Return(test1Xxx.Props(), nil)
			mxg.EXPECT().Update(gomock.Any(), gomock.Any()).Return(test1Xxx.Props(), nil)
		},
		Ctx: context.Background(),
		Args: &xxx.UsecaseUpdateInput{
			Id:     test1Xxx.Id(),
			Name:   test1XxxUpdated.Name(),
			Number: test1XxxUpdated.Number(),
		},
		Expected: &xxx.UsecaseUpdateOutput{
			Xxx: test1XxxUpdated.Props(),
		},
		IsErr: false,
	}
	tests = append(tests, test1)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Setup()

			usecase := xi.NewXxxInteractorUpdate(mxg)

			result, err := usecase.Handle(test.Ctx, test.Args.(*xxx.UsecaseUpdateInput))

			if test.IsErr && err != nil {
				assert.Equal(t, test.Err.(*dErr.Error).Type, err.(*dErr.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.Expected.(*xxx.UsecaseUpdateOutput).Xxx.Id(), result.Xxx.Id())
			assert.Equal(t, test.Expected.(*xxx.UsecaseUpdateOutput).Xxx.Name(), result.Xxx.Name())
			assert.Equal(t, test.Expected.(*xxx.UsecaseUpdateOutput).Xxx.Number(), result.Xxx.Number())
			assert.Equal(t, test.Expected.(*xxx.UsecaseUpdateOutput).Xxx.CreatedAt(), result.Xxx.CreatedAt())
			assert.NotEqual(t, test.Expected.(*xxx.UsecaseUpdateOutput).Xxx.UpdatedAt(), result.Xxx.UpdatedAt())
		})
	}
}
