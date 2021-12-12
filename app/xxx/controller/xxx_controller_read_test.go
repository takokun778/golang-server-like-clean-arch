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

func TestUserControllerRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mxu := mx.NewMockUsecase(ctrl)

	tests := make([]test.Case, 0)

	test1Xxx := test.NewXxx("test1", 1)
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() interface{} {
			input := &xxx.UsecaseReadInput{
				Id: test1Xxx.Id(),
			}
			output := &xxx.UsecaseReadOutput{
				Xxx: test1Xxx,
			}
			mxu.EXPECT().Read(gomock.Any(), input).Return(output, nil)
			return nil
		},
		Ctx: context.Background(),
		Args: &pbXxx.ReadRequest{
			Id: test1Xxx.Id().Value().String(),
		},
		Expected: &pbXxx.ReadResponse{
			Xxx: &pbXxx.Xxx{
				Id:        test1Xxx.Id().Value().String(),
				Name:      test1Xxx.Name().Value(),
				Number:    int32(test1Xxx.Number().Value()),
				CreatedAt: timestamppb.New(test1Xxx.CreatedAt().Value()),
				UpdatedAt: timestamppb.New(test1Xxx.UpdatedAt().Value()),
			},
		},
		IsErr: false,
	}
	tests = append(tests, test1)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Setup()

			controller := mc.NewXxxController(mxu)

			result, err := controller.Read(test.Ctx, test.Args.(*pbXxx.ReadRequest))

			if test.IsErr && err != nil {
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.Expected.(*pbXxx.ReadResponse), result)
		})
	}
}
