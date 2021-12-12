package xxx_test

import (
	"context"
	"fmt"
	"testing"

	pbXxx "xxx/proto/xxx"
	"xxx/test"
	"xxx/test/xxx"

	"github.com/stretchr/testify/assert"
)

func TestXxxRead(t *testing.T) {
	client := xxx.ClientConnect()

	tests := make([]test.Case, 0)

	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() interface{} {
			result, _ := client.Create(context.Background(), &pbXxx.CreateRequest{
				Name:   "test1",
				Number: 1,
			})
			return result.Xxx.Id
		},
		Ctx:  context.Background(),
		Args: &pbXxx.ReadRequest{},
		Expected: &pbXxx.ReadResponse{
			Xxx: &pbXxx.Xxx{
				Name:   "test1",
				Number: 1,
			},
		},
		IsErr: false,
	}
	tests = append(tests, test1)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			request := tt.Args.(*pbXxx.ReadRequest)

			request.Id = tt.Setup().(string)

			result, err := client.Read(tt.Ctx, request)

			if tt.IsErr && err != nil {
				return
			} else {
				fmt.Println(err)
				assert.NoError(t, err)
			}

			assert.Equal(t, request.Id, result.Xxx.Id)
			assert.Equal(t, tt.Expected.(*pbXxx.ReadResponse).Xxx.Name, result.Xxx.Name)
			assert.Equal(t, tt.Expected.(*pbXxx.ReadResponse).Xxx.Number, result.Xxx.Number)
		})
	}
}
