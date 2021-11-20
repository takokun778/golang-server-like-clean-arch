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

func TestUserControllerReadAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mxu := mx.NewMockUsecase(ctrl)

	tests := make([]test.Case, 0)

	test1XxxValues1 := xxx.Create(xxx.Name("test1"), xxx.Number(1)).Values()
	test1XxxValues2 := xxx.Create(xxx.Name("test1"), xxx.Number(2)).Values()
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			input := &xxx.UsecaseReadAllInput{}
			output := &xxx.UsecaseReadAllOutput{
				Xxxs: []xxx.Values{test1XxxValues1, test1XxxValues2},
			}
			mxu.EXPECT().ReadAll(gomock.Any(), input).Return(output, nil)
		},
		Ctx:  context.Background(),
		Args: &pbXxx.ReadAllRequest{},
		Expected: &pbXxx.ReadAllResponse{
			Xxxs: []*pbXxx.Xxx{
				{
					Id:        test1XxxValues1.Id().Value().String(),
					Name:      test1XxxValues1.Name().Value(),
					Number:    int32(test1XxxValues1.Number().Value()),
					CreatedAt: timestamppb.New(test1XxxValues1.CreatedAt().Value()),
					UpdatedAt: timestamppb.New(test1XxxValues1.UpdatedAt().Value()),
				},
				{
					Id:        test1XxxValues2.Id().Value().String(),
					Name:      test1XxxValues2.Name().Value(),
					Number:    int32(test1XxxValues2.Number().Value()),
					CreatedAt: timestamppb.New(test1XxxValues2.CreatedAt().Value()),
					UpdatedAt: timestamppb.New(test1XxxValues2.UpdatedAt().Value()),
				},
			},
		},
		IsErr: false,
	}
	tests = append(tests, test1)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Setup()

			controller := mc.NewXxxController(mxu)

			result, err := controller.ReadAll(test.Ctx, test.Args.(*pbXxx.ReadAllRequest))

			if test.IsErr && err != nil {
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.Expected.(*pbXxx.ReadAllResponse), result)
		})
	}
}
