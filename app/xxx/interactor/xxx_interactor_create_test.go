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

func TestXxxInteractorCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mxg := mx.NewMockRepository(ctrl)

	tests := make([]test.Case, 0)

	test1Xxx := test.NewXxx("test1", 1).Props()
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			mxg.EXPECT().Save(gomock.Any(), gomock.Any()).Return(test1Xxx, nil)
		},
		Ctx: context.Background(),
		Args: &xxx.UsecaseCreateInput{
			Name:   test1Xxx.Name(),
			Number: test1Xxx.Number(),
		},
		Expected: &xxx.UsecaseCreateOutput{
			Xxx: test1Xxx,
		},
		IsErr: false,
	}
	tests = append(tests, test1)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Setup()

			usecase := xi.NewXxxInteractor(mxg)

			result, err := usecase.Create(test.Ctx, test.Args.(*xxx.UsecaseCreateInput))

			if test.IsErr && err != nil {
				assert.Equal(t, test.Err.(*dErr.Error).Type, err.(*dErr.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.Expected.(*xxx.UsecaseCreateOutput).Xxx.Name(), result.Xxx.Name())
			assert.Equal(t, test.Expected.(*xxx.UsecaseCreateOutput).Xxx.Number(), result.Xxx.Number())
		})
	}
}
