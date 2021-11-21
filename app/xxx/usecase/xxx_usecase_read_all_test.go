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

func TestXxxUsecaseReadAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mxg := mx.NewMockRepository(ctrl)

	tests := make([]test.Case, 0)

	test1Xxx1 := xxx.Create(xxx.Name("test1"), xxx.Number(1))
	test1Xxx2 := xxx.Create(xxx.Name("test1"), xxx.Number(2))
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			mxg.EXPECT().FindAll(gomock.Any()).Return([]xxx.Props{test1Xxx1.Props(), test1Xxx2.Props()}, nil)
		},
		Ctx:  context.Background(),
		Args: &xxx.UsecaseReadAllInput{},
		Expected: &xxx.UsecaseReadAllOutput{
			Xxxs: []xxx.Props{test1Xxx1.Props(), test1Xxx2.Props()},
		},
		IsErr: false,
	}
	tests = append(tests, test1)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Setup()

			usecase := xu.NewXxxUsecase(mxg)

			result, err := usecase.ReadAll(test.Ctx, test.Args.(*xxx.UsecaseReadAllInput))

			if test.IsErr && err != nil {
				assert.Equal(t, test.Err.(*common.Error).Type, err.(*common.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.Expected.(*xxx.UsecaseReadAllOutput).Xxxs, result.Xxxs)
		})
	}
}
