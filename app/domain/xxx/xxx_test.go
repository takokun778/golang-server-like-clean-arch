package xxx_test

import (
	"context"
	"testing"

	"xxx/app/domain/common"
	"xxx/app/domain/xxx"
	"xxx/test"

	"github.com/stretchr/testify/assert"
)

func TestXxxReconstruct(t *testing.T) {
	tests := make([]test.Case, 0)

	test1Args := xxx.NewValues(
		common.CreateRandomId(),
		xxx.Name("test1"),
		xxx.Number(1),
		common.Now(),
		common.Now(),
	)
	test1Expected := xxx.NewValues(
		test1Args.Id(),
		test1Args.Name(),
		test1Args.Number(),
		test1Args.CreatedAt(),
		test1Args.UpdatedAt(),
	)
	test1 := test.Case{
		Name:     "正常動作確認",
		Setup:    func() {},
		Ctx:      context.Background(),
		Args:     test1Args,
		Expected: test1Expected,
		IsErr:    false,
	}
	tests = append(tests, test1)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Setup()

			result := xxx.Reconstruct(
				tt.Args.(xxx.Values),
			)

			assert.Equal(t, tt.Expected.(xxx.Values).Name(), result.Name())
			assert.Equal(t, tt.Expected.(xxx.Values).Number(), result.Number())
		})
	}
}

func TestXxxCreate(t *testing.T) {
	tests := make([]test.Case, 0)

	test1Values := xxx.NewValues(
		common.CreateRandomId(),
		xxx.Name("test1"),
		xxx.Number(1),
		common.Now(),
		common.Now(),
	)
	test1 := test.Case{
		Name:     "正常動作確認",
		Setup:    func() {},
		Ctx:      context.Background(),
		Args:     test1Values,
		Expected: test1Values,
		IsErr:    false,
	}
	tests = append(tests, test1)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Setup()

			result := xxx.Create(
				tt.Args.(xxx.Values).Name(),
				tt.Args.(xxx.Values).Number(),
			)

			assert.Equal(t, tt.Expected.(xxx.Values).Name(), result.Name())
			assert.Equal(t, tt.Expected.(xxx.Values).Number(), result.Number())
		})
	}
}

func TestXxxUpdate(t *testing.T) {
	tests := make([]test.Case, 0)

	test1Args := xxx.NewValues(
		common.CreateRandomId(),
		xxx.Name("updated"),
		xxx.Number(2),
		common.Now(),
		common.Now(),
	)
	test1Expected := xxx.NewValues(
		common.CreateRandomId(),
		xxx.Name("updated"),
		xxx.Number(2),
		common.Now(),
		common.Now(),
	)
	test1 := test.Case{
		Name:     "正常動作確認",
		Setup:    func() {},
		Ctx:      context.Background(),
		Args:     test1Args,
		Expected: test1Expected,
		IsErr:    false,
	}
	tests = append(tests, test1)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Setup()

			src := xxx.Create(
				xxx.Name("create"),
				xxx.Number(1),
			)

			result := src.Update(
				tt.Args.(xxx.Values).Name(),
				tt.Args.(xxx.Values).Number(),
			)

			assert.Equal(t, tt.Expected.(xxx.Values).Name(), result.Name())
			assert.Equal(t, tt.Expected.(xxx.Values).Number(), result.Number())
		})
	}
}
