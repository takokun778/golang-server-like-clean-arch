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

func TestUserControllerUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mxu := mx.NewMockUsecase(ctrl)

	tests := make([]test.Case, 0)

	test1Xxx := test.NewXxx("test1", 1).Props()
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			input := &xxx.UsecaseUpdateInput{
				Id:     test1Xxx.Id(),
				Name:   test1Xxx.Name(),
				Number: test1Xxx.Number(),
			}
			output := &xxx.UsecaseUpdateOutput{
				Xxx: test1Xxx,
			}
			mxu.EXPECT().Update(gomock.Any(), input).Return(output, nil)
		},
		Ctx: context.Background(),
		Args: &pbXxx.UpdateRequest{
			Id:     test1Xxx.Id().Value().String(),
			Name:   test1Xxx.Name().Value(),
			Number: int32(test1Xxx.Number().Value()),
		},
		Expected: &pbXxx.UpdateResponse{
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

			result, err := controller.Update(test.Ctx, test.Args.(*pbXxx.UpdateRequest))

			if test.IsErr && err != nil {
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.Expected.(*pbXxx.UpdateResponse), result)
		})
	}
}
