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

	test1XxxValues := xxx.Create(xxx.Name("test1"), xxx.Number(1)).Values()
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			input := &xxx.UsecaseUpdateInput{
				Id:     test1XxxValues.Id(),
				Name:   test1XxxValues.Name(),
				Number: test1XxxValues.Number(),
			}
			output := &xxx.UsecaseUpdateOutput{
				Xxx: test1XxxValues,
			}
			mxu.EXPECT().Update(gomock.Any(), input).Return(output, nil)
		},
		Ctx: context.Background(),
		Args: &pbXxx.UpdateRequest{
			Id:     test1XxxValues.Id().Value().String(),
			Name:   test1XxxValues.Name().Value(),
			Number: int32(test1XxxValues.Number().Value()),
		},
		Expected: &pbXxx.UpdateResponse{
			Xxx: &pbXxx.Xxx{
				Id:        test1XxxValues.Id().Value().String(),
				Name:      test1XxxValues.Name().Value(),
				Number:    int32(test1XxxValues.Number().Value()),
				CreatedAt: timestamppb.New(test1XxxValues.CreatedAt().Value()),
				UpdatedAt: timestamppb.New(test1XxxValues.UpdatedAt().Value()),
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
