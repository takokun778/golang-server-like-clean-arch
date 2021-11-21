package controller_test

import (
	"context"
	"testing"

	"xxx/app/domain/xxx"
	mc "xxx/app/xxx/controller"
	mx "xxx/mock/xxx"
	pbXxx "xxx/proto/xxx"

	"xxx/test"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestUserControllerDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mxu := mx.NewMockUsecase(ctrl)

	tests := make([]test.Case, 0)

	test1XxxProps := xxx.Create(xxx.Name("test1"), xxx.Number(1)).Props()
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			input := &xxx.UsecaseDeleteInput{
				Id: test1XxxProps.Id(),
			}
			output := &xxx.UsecaseDeleteOutput{
				Xxx: test1XxxProps,
			}
			mxu.EXPECT().Delete(gomock.Any(), input).Return(output, nil)
		},
		Ctx: context.Background(),
		Args: &pbXxx.DeleteRequest{
			Id: test1XxxProps.Id().Value().String(),
		},
		Expected: &pbXxx.DeleteResponse{
			Xxx: &pbXxx.Xxx{
				Id:        test1XxxProps.Id().Value().String(),
				Name:      test1XxxProps.Name().Value(),
				Number:    int32(test1XxxProps.Number().Value()),
				CreatedAt: timestamppb.New(test1XxxProps.CreatedAt().Value()),
				UpdatedAt: timestamppb.New(test1XxxProps.UpdatedAt().Value()),
			},
		},
		IsErr: false,
	}
	tests = append(tests, test1)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Setup()

			controller := mc.NewXxxController(mxu)

			result, err := controller.Delete(test.Ctx, test.Args.(*pbXxx.DeleteRequest))

			if test.IsErr && err != nil {
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.Expected.(*pbXxx.DeleteResponse), result)
		})
	}
}
