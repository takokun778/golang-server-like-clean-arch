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

func TestXxxUsecaseUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mxg := mx.NewMockRepository(ctrl)

	tests := make([]test.Case, 0)

	test1Xxx := xxx.Create(xxx.Name("test1"), xxx.Number(1))
	test1XxxUpdated := test1Xxx.Update(xxx.Name("updated1"), xxx.Number(2))
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			mxg.EXPECT().Find(gomock.Any(), test1Xxx.Id()).Return(test1Xxx.Props(), nil)
			mxg.EXPECT().Update(gomock.Any(), gomock.Any()).Return(test1XxxUpdated.Props(), nil)
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

			usecase := xu.NewXxxUsecase(mxg)

			result, err := usecase.Update(test.Ctx, test.Args.(*xxx.UsecaseUpdateInput))

			if test.IsErr && err != nil {
				assert.Equal(t, test.Err.(*common.Error).Type, err.(*common.Error).Type)
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
