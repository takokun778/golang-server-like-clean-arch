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

func TestXxxUsecaseDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mxg := mx.NewMockRepository(ctrl)

	tests := make([]test.Case, 0)

	test1Xxx := test.NewXxx("test1", 1).Props()
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			mxg.EXPECT().Find(gomock.Any(), &xxx.RepositoryFindItem{Id: test1Xxx.Id()}).Return(test1Xxx, nil)
			mxg.EXPECT().Delete(gomock.Any(), &xxx.RepositoryDeleteItem{Id: test1Xxx.Id()}).Return(nil)
		},
		Ctx: context.Background(),
		Args: &xxx.UsecaseDeleteInput{
			Id: test1Xxx.Id(),
		},
		Expected: &xxx.UsecaseDeleteOutput{
			Xxx: test1Xxx,
		},
		IsErr: false,
	}
	tests = append(tests, test1)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Setup()

			usecase := xu.NewXxxUsecase(mxg)

			result, err := usecase.Delete(test.Ctx, test.Args.(*xxx.UsecaseDeleteInput))

			if test.IsErr && err != nil {
				assert.Equal(t, test.Err.(*xe.Error).Type, err.(*xe.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.Expected.(*xxx.UsecaseDeleteOutput).Xxx, result.Xxx)
		})
	}
}
