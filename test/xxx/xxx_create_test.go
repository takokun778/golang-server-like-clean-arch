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

func TestXxxCreate(t *testing.T) {
	client := xxx.ClientConnect()

	tests := make([]test.Case, 0)

	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() interface{} {
			return nil
		},
		Ctx: context.Background(),
		Args: &pbXxx.CreateRequest{
			Name:   "test1",
			Number: 1,
		},
		Expected: &pbXxx.CreateResponse{
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
			tt.Setup()

			result, err := client.Create(tt.Ctx, tt.Args.(*pbXxx.CreateRequest))

			if tt.IsErr && err != nil {
				return
			} else {
				fmt.Println(err)
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.Expected.(*pbXxx.CreateResponse).Xxx.Name, result.Xxx.Name)
			assert.Equal(t, tt.Expected.(*pbXxx.CreateResponse).Xxx.Number, result.Xxx.Number)
		})
	}
}
