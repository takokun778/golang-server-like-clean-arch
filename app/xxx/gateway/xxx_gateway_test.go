package gateway_test

import (
	"context"
	"errors"
	"testing"

	cg "xxx/app/common/gateway"
	xe "xxx/app/domain/error"
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

	test1Xxx := test.NewXxx("test1", 1).Props()
	test1 := test.Case{
		Name:  "正常動作確認",
		Setup: func() {},
		Ctx:   context.Background(),
		Args: &xxx.RepositorySaveInput{
			Xxx: test1Xxx,
		},
		Expected: &xxx.RepositorySaveOutput{
			Xxx: test1Xxx,
		},
		IsErr: false,
	}
	tests = append(tests, test1)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Setup()

			result, err := gateway.Save(tt.Ctx, tt.Args.(*xxx.RepositorySaveInput))

			if tt.IsErr && err != nil {
				assert.Equal(t, tt.Err.(*xe.Error).Type, err.(*xe.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.Expected.(*xxx.RepositorySaveOutput), result)
		})
	}
}

func TestXxxGatewayFind(t *testing.T) {
	tests := make([]test.Case, 0)

	test1Xxx := test.NewXxx("test1", 1).Props()
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			gateway.Save(context.Background(), &xxx.RepositorySaveInput{Xxx: test1Xxx})
		},
		Ctx: context.Background(),
		Args: &xxx.RepositoryFindInput{
			Id: test1Xxx.Id(),
		},
		Expected: &xxx.RepositoryFindOutput{
			Xxx: test1Xxx,
		},
		IsErr: false,
	}
	tests = append(tests, test1)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Setup()

			result, err := gateway.Find(tt.Ctx, tt.Args.(*xxx.RepositoryFindInput))

			if tt.IsErr && err != nil {
				assert.Equal(t, tt.Err.(*xe.Error).Type, err.(*xe.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.Expected.(*xxx.RepositoryFindOutput), result)
		})
	}
}

func TestXxxGatewayFindAll(t *testing.T) {
	tests := make([]test.Case, 0)

	test1Xxx1 := test.NewXxx("test1", 1).Props()
	test1Xxx2 := test.NewXxx("test1", 1).Props()
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			resetTable()
			gateway.Save(context.Background(), &xxx.RepositorySaveInput{Xxx: test1Xxx1})
			gateway.Save(context.Background(), &xxx.RepositorySaveInput{Xxx: test1Xxx2})
		},
		Ctx:  context.Background(),
		Args: &xxx.RepositoryFindAllInput{},
		Expected: &xxx.RepositoryFindAllOutput{
			Xxxs: []xxx.Props{test1Xxx1, test1Xxx2},
		},
		IsErr: false,
	}
	tests = append(tests, test1)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Setup()

			result, err := gateway.FindAll(tt.Ctx, tt.Args.(*xxx.RepositoryFindAllInput))

			if tt.IsErr && err != nil {
				assert.Equal(t, tt.Err.(*xe.Error).Type, err.(*xe.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.Expected.(*xxx.RepositoryFindAllOutput), result)
		})
	}
}

func TestXxxGatewayUpdate(t *testing.T) {
	tests := make([]test.Case, 0)

	test1Xxx := test.NewXxx("test1", 1).Props()
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			gateway.Save(context.Background(), &xxx.RepositorySaveInput{Xxx: test1Xxx})
		},
		Ctx: context.Background(),
		Args: &xxx.RepositoryUpdateInput{
			Xxx: test1Xxx,
		},
		Expected: &xxx.RepositoryUpdateOutput{
			Xxx: test1Xxx,
		},
		IsErr: false,
	}
	tests = append(tests, test1)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Setup()

			result, err := gateway.Update(tt.Ctx, tt.Args.(*xxx.RepositoryUpdateInput))

			if tt.IsErr && err != nil {
				assert.Equal(t, tt.Err.(*xe.Error).Type, err.(*xe.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.Expected.(*xxx.RepositoryUpdateOutput), result)
		})
	}
}

func TestXxxGatewayDelete(t *testing.T) {
	tests := make([]test.Case, 0)

	test1Xxx := test.NewXxx("test1", 1).Props()
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			gateway.Save(context.Background(), &xxx.RepositorySaveInput{Xxx: test1Xxx})
		},
		Ctx: context.Background(),
		Args: &xxx.RepositoryDeleteInput{
			Id: test1Xxx.Id(),
		},
		Expected: &xxx.RepositoryDeleteOutput{},
		IsErr:    false,
	}
	tests = append(tests, test1)

	test2 := test.Case{
		Name:  "存在しないデータを削除",
		Setup: func() {},
		Ctx:   context.Background(),
		Args: &xxx.RepositoryDeleteInput{
			Id: test1Xxx.Id(),
		},
		Expected: &xxx.RepositoryDeleteOutput{},
		IsErr:    true,
		Err:      xe.NewInternalServerError(errors.New(""), ""),
	}
	tests = append(tests, test2)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Setup()

			_, err := gateway.Delete(tt.Ctx, tt.Args.(*xxx.RepositoryDeleteInput))

			if tt.IsErr && err != nil {
				assert.Equal(t, tt.Err.(*xe.Error).Type, err.(*xe.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			_, err = gateway.Find(tt.Ctx, &xxx.RepositoryFindInput{Id: tt.Args.(*xxx.RepositoryDeleteInput).Id})

			if err == nil {
				t.Fail()
			}

		})
	}
}
