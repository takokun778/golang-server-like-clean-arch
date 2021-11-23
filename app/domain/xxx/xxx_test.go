package xxx_test

import (
	"context"
	"testing"

	"xxx/app/domain/xxx"
	"xxx/test"

	"github.com/stretchr/testify/assert"
)

func TestXxxReconstruct(t *testing.T) {
	tests := make([]test.Case, 0)

	test1Name, _ := xxx.NewName("test1")
	test1Number, _ := xxx.NewNumber(1)
	test1Args := xxx.NewProps(
		xxx.CreateRandomId(),
		test1Name,
		test1Number,
		xxx.TimeNow(),
		xxx.TimeNow(),
	)
	test1Expected := xxx.NewProps(
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
				tt.Args.(xxx.Props),
			)

			assert.Equal(t, tt.Expected.(xxx.Props).Name(), result.Name())
			assert.Equal(t, tt.Expected.(xxx.Props).Number(), result.Number())
		})
	}
}

func TestXxxCreate(t *testing.T) {
	tests := make([]test.Case, 0)

	test1Name, _ := xxx.NewName("test1")
	test1Number, _ := xxx.NewNumber(1)
	test1Props := xxx.NewProps(
		xxx.CreateRandomId(),
		test1Name,
		test1Number,
		xxx.TimeNow(),
		xxx.TimeNow(),
	)
	test1 := test.Case{
		Name:     "正常動作確認",
		Setup:    func() {},
		Ctx:      context.Background(),
		Args:     test1Props,
		Expected: test1Props,
		IsErr:    false,
	}
	tests = append(tests, test1)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Setup()

			result := xxx.Create(
				tt.Args.(xxx.Props).Name(),
				tt.Args.(xxx.Props).Number(),
			)

			assert.Equal(t, tt.Expected.(xxx.Props).Name(), result.Name())
			assert.Equal(t, tt.Expected.(xxx.Props).Number(), result.Number())
		})
	}
}

func TestXxxUpdate(t *testing.T) {
	tests := make([]test.Case, 0)

	test1Name, _ := xxx.NewName("test1")
	test1Number, _ := xxx.NewNumber(1)
	test1Args := xxx.NewProps(
		xxx.CreateRandomId(),
		test1Name,
		test1Number,
		xxx.TimeNow(),
		xxx.TimeNow(),
	)
	test1Expected := xxx.NewProps(
		xxx.CreateRandomId(),
		test1Name,
		test1Number,
		xxx.TimeNow(),
		xxx.TimeNow(),
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

			name, _ := xxx.NewName("create")
			number, _ := xxx.NewNumber(1)
			src := xxx.Create(
				name,
				number,
			)

			result := src.Update(
				tt.Args.(xxx.Props).Name(),
				tt.Args.(xxx.Props).Number(),
			)

			assert.Equal(t, tt.Expected.(xxx.Props).Name(), result.Name())
			assert.Equal(t, tt.Expected.(xxx.Props).Number(), result.Number())
		})
	}
}
