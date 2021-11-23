package usecase_test

import (
	"context"
	"testing"

	xe "xxx/app/domain/error"
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

	test1Xxx1 := test.NewXxx("test1", 1).Props()
	test1Xxx2 := test.NewXxx("test1", 1).Props()
	test1Output := &xxx.RepositoryFindAllOutput{
		Xxxs: []xxx.Props{test1Xxx1, test1Xxx2},
	}
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			mxg.EXPECT().FindAll(gomock.Any(), &xxx.RepositoryFindAllInput{}).Return(test1Output, nil)
		},
		Ctx:  context.Background(),
		Args: &xxx.UsecaseReadAllInput{},
		Expected: &xxx.UsecaseReadAllOutput{
			Xxxs: []xxx.Props{test1Xxx1, test1Xxx2},
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
				assert.Equal(t, test.Err.(*xe.Error).Type, err.(*xe.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.Expected.(*xxx.UsecaseReadAllOutput).Xxxs, result.Xxxs)
		})
	}
}
