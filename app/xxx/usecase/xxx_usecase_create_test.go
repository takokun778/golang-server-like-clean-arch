package usecase_test

import (
	"context"
	"testing"

	"xxx/app/domain/common"
	"xxx/app/domain/xxx"
	xu "xxx/app/xxx/usecase"
	mx "xxx/mock/xxx"
	"xxx/test"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestXxxUsecaseCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mxg := mx.NewMockRepository(ctrl)

	tests := make([]test.Case, 0)

	test1Xxx := xxx.Create(xxx.Name("test1"), xxx.Number(1))
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			mxg.EXPECT().Save(gomock.Any(), gomock.Any()).Return(test1Xxx.Values(), nil)
		},
		Ctx: context.Background(),
		Args: &xxx.UsecaseCreateInput{
			Name:   test1Xxx.Name(),
			Number: test1Xxx.Number(),
		},
		Expected: &xxx.UsecaseCreateOutput{
			Xxx: test1Xxx.Values(),
		},
		IsErr: false,
	}
	tests = append(tests, test1)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Setup()

			usecase := xu.NewXxxUsecase(mxg)

			result, err := usecase.Create(test.Ctx, test.Args.(*xxx.UsecaseCreateInput))

			if test.IsErr && err != nil {
				assert.Equal(t, test.Err.(*common.Error).Type, err.(*common.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.Expected.(*xxx.UsecaseCreateOutput).Xxx.Name(), result.Xxx.Name())
			assert.Equal(t, test.Expected.(*xxx.UsecaseCreateOutput).Xxx.Number(), result.Xxx.Number())
		})
	}
}
