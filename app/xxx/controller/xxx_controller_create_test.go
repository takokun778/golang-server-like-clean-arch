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

func TestUserControllerCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mxuc := mx.NewMockUsecaseCreate(ctrl)
	mxur := mx.NewMockUsecaseRead(ctrl)
	mxura := mx.NewMockUsecaseReadAll(ctrl)
	mxuu := mx.NewMockUsecaseUpdate(ctrl)
	mxud := mx.NewMockUsecaseDelete(ctrl)

	tests := make([]test.Case, 0)

	test1Xxx := test.NewXxx("test1", 1).Props()
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			input := &xxx.UsecaseCreateInput{
				Name:   test1Xxx.Name(),
				Number: test1Xxx.Number(),
			}
			output := &xxx.UsecaseCreateOutput{
				Xxx: test1Xxx,
			}
			mxuc.EXPECT().Handle(gomock.Any(), input).Return(output, nil)
		},
		Ctx: context.Background(),
		Args: &pbXxx.CreateRequest{
			Name:   "test1",
			Number: 1,
		},
		Expected: &pbXxx.CreateResponse{
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

			controller := mc.NewXxxController(mxuc, mxur, mxura, mxuu, mxud)

			result, err := controller.Create(test.Ctx, test.Args.(*pbXxx.CreateRequest))

			if test.IsErr && err != nil {
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.Expected.(*pbXxx.CreateResponse), result)
		})
	}
}
