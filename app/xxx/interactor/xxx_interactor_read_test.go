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

func TestXxxInteractorRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mxg := mx.NewMockRepository(ctrl)

	tests := make([]test.Case, 0)

	test1Xxx := test.NewXxx("test1", 1)
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() interface{} {
			mxg.EXPECT().Find(gomock.Any(), &xxx.RepositoryFindItem{Id: test1Xxx.Id()}).Return(test1Xxx, nil)
			return nil
		},
		Ctx: context.Background(),
		Args: &xxx.UsecaseReadInput{
			Id: test1Xxx.Id(),
		},
		Expected: &xxx.UsecaseReadOutput{
			Xxx: test1Xxx,
		},
		IsErr: false,
	}
	tests = append(tests, test1)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Setup()

			usecase := xi.NewXxxInteractor(mxg)

			result, err := usecase.Read(test.Ctx, test.Args.(*xxx.UsecaseReadInput))

			if test.IsErr && err != nil {
				assert.Equal(t, test.Err.(*dErr.Error).Type, err.(*dErr.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.Expected.(*xxx.UsecaseReadOutput).Xxx, result.Xxx)
		})
	}
}
