package gateway_test

import (
	"context"
	"errors"
	"testing"

	cg "xxx/app/common/gateway"
	"xxx/app/domain/common"
	"xxx/app/domain/xxx"
	xg "xxx/app/xxx/gateway"
	"xxx/test"

	"github.com/stretchr/testify/assert"
)

var gateway = xg.NewXxxGateway()

func TestMain(m *testing.M) {
	// テスト前処理
	resetTable()

	// テスト実行
	m.Run()

	// テスト後処理
	resetTable()
}

func resetTable() {
	database := cg.DatabaseConnect()

	ctx := context.Background()

	database.Xxx.Delete().Exec(ctx)
}

func TestXxxGatewaySave(t *testing.T) {
	tests := make([]test.Case, 0)

	test1Value := xxx.Create(
		xxx.Name("test1"),
		xxx.Number(1),
	).Values()
	test1 := test.Case{
		Name:     "正常動作確認",
		Setup:    func() {},
		Ctx:      context.Background(),
		Args:     test1Value,
		Expected: test1Value,
		IsErr:    false,
	}
	tests = append(tests, test1)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Setup()

			result, err := gateway.Save(tt.Ctx, tt.Args.(xxx.Values))

			if tt.IsErr && err != nil {
				assert.Equal(t, tt.Err.(*common.Error).Type, err.(*common.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.Expected.(xxx.Values), result)
		})
	}
}

func TestXxxGatewayFind(t *testing.T) {
	tests := make([]test.Case, 0)

	test1Value := xxx.Create(
		xxx.Name("test1"),
		xxx.Number(1),
	).Values()
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			gateway.Save(context.Background(), test1Value)
		},
		Ctx:      context.Background(),
		Args:     test1Value.Id(),
		Expected: test1Value,
		IsErr:    false,
	}
	tests = append(tests, test1)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Setup()

			result, err := gateway.Find(tt.Ctx, tt.Args.(common.Id))

			if tt.IsErr && err != nil {
				assert.Equal(t, tt.Err.(*common.Error).Type, err.(*common.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.Expected.(xxx.Values), result)
		})
	}
}

func TestXxxGatewayFindAll(t *testing.T) {
	tests := make([]test.Case, 0)

	test1Value1 := xxx.Create(
		xxx.Name("test1"),
		xxx.Number(1),
	).Values()
	test1Value2 := xxx.Create(
		xxx.Name("test1"),
		xxx.Number(1),
	).Values()

	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			resetTable()
			gateway.Save(context.Background(), test1Value1)
			gateway.Save(context.Background(), test1Value2)
		},
		Ctx:      context.Background(),
		Args:     nil,
		Expected: []xxx.Values{test1Value1, test1Value2},
		IsErr:    false,
	}
	tests = append(tests, test1)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Setup()

			result, err := gateway.FindAll(tt.Ctx)

			if tt.IsErr && err != nil {
				assert.Equal(t, tt.Err.(*common.Error).Type, err.(*common.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.Expected.([]xxx.Values), result)
		})
	}
}

func TestXxxGatewayUpdate(t *testing.T) {
	tests := make([]test.Case, 0)

	test1Value := xxx.Create(
		xxx.Name("test1"),
		xxx.Number(1),
	).Values()
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			gateway.Save(context.Background(), test1Value)
		},
		Ctx:      context.Background(),
		Args:     test1Value,
		Expected: test1Value,
		IsErr:    false,
	}
	tests = append(tests, test1)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Setup()

			result, err := gateway.Update(tt.Ctx, tt.Args.(xxx.Values))

			if tt.IsErr && err != nil {
				assert.Equal(t, tt.Err.(*common.Error).Type, err.(*common.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.Expected.(xxx.Values), result)
		})
	}
}

func TestXxxGatewayDelete(t *testing.T) {
	tests := make([]test.Case, 0)

	test1Value := xxx.Create(
		xxx.Name("test1"),
		xxx.Number(1),
	).Values()
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			gateway.Save(context.Background(), test1Value)
		},
		Ctx:      context.Background(),
		Args:     test1Value.Id(),
		Expected: test1Value,
		IsErr:    false,
	}
	tests = append(tests, test1)

	test2 := test.Case{
		Name:     "存在しないデータを削除",
		Setup:    func() {},
		Ctx:      context.Background(),
		Args:     test1Value.Id(),
		Expected: test1Value,
		IsErr:    true,
		Err:      common.NewInternalServerError(errors.New(""), ""),
	}
	tests = append(tests, test2)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Setup()

			err := gateway.Delete(tt.Ctx, tt.Args.(common.Id))

			if tt.IsErr && err != nil {
				assert.Equal(t, tt.Err.(*common.Error).Type, err.(*common.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			_, err = gateway.Find(tt.Ctx, tt.Args.(common.Id))

			if err == nil {
				t.Fail()
			}

		})
	}
}
