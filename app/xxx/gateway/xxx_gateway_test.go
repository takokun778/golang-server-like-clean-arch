package gateway_test

import (
	"context"
	"errors"
	"testing"

	dErr "xxx/app/domain/error"
	"xxx/app/domain/xxx"
	"xxx/app/infra/ent"
	xg "xxx/app/xxx/gateway"
	"xxx/test"

	"github.com/stretchr/testify/assert"
)

var gateway = xg.NewXxxGateway(ent.DatabaseConnect())

func TestMain(m *testing.M) {
	// テスト前処理
	resetTable()

	// テスト実行
	m.Run()

	// テスト後処理
	resetTable()
}

func resetTable() {
	database := ent.DatabaseConnect()

	ctx := context.Background()

	database.Xxx.Delete().Exec(ctx)
}

func TestXxxGatewaySave(t *testing.T) {
	tests := make([]test.Case, 0)

	test1Xxx := test.NewXxx("test1", 1)
	test1 := test.Case{
		Name:  "正常動作確認",
		Setup: func() {},
		Ctx:   context.Background(),
		Args: &xxx.RepositorySaveItem{
			Xxx: test1Xxx,
		},
		Expected: test1Xxx,
		IsErr:    false,
	}
	tests = append(tests, test1)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Setup()

			result, err := gateway.Save(tt.Ctx, tt.Args.(*xxx.RepositorySaveItem))

			if tt.IsErr && err != nil {
				assert.Equal(t, tt.Err.(*dErr.Error).Type, err.(*dErr.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.Expected.(xxx.Props), result)
		})
	}
}

func TestXxxGatewayFind(t *testing.T) {
	tests := make([]test.Case, 0)

	test1Xxx := test.NewXxx("test1", 1)
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			gateway.Save(context.Background(), &xxx.RepositorySaveItem{Xxx: test1Xxx})
		},
		Ctx: context.Background(),
		Args: &xxx.RepositoryFindItem{
			Id: test1Xxx.Id(),
		},
		Expected: test1Xxx,
		IsErr:    false,
	}
	tests = append(tests, test1)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Setup()

			result, err := gateway.Find(tt.Ctx, tt.Args.(*xxx.RepositoryFindItem))

			if tt.IsErr && err != nil {
				assert.Equal(t, tt.Err.(*dErr.Error).Type, err.(*dErr.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.Expected.(xxx.Props).Id().Value(), result.Id().Value())
			assert.Equal(t, tt.Expected.(xxx.Props).Name().Value(), result.Name().Value())
			assert.Equal(t, tt.Expected.(xxx.Props).Number().Value(), result.Number().Value())
		})
	}
}

func TestXxxGatewayFindAll(t *testing.T) {
	tests := make([]test.Case, 0)

	test1Xxx1 := test.NewXxx("test1", 1)
	test1Xxx2 := test.NewXxx("test1", 1)
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			resetTable()
			gateway.Save(context.Background(), &xxx.RepositorySaveItem{Xxx: test1Xxx1})
			gateway.Save(context.Background(), &xxx.RepositorySaveItem{Xxx: test1Xxx2})
		},
		Ctx:      context.Background(),
		Args:     &xxx.RepositoryFindAllItem{},
		Expected: []xxx.Props{test1Xxx1, test1Xxx2},
		IsErr:    false,
	}
	tests = append(tests, test1)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Setup()

			result, err := gateway.FindAll(tt.Ctx, tt.Args.(*xxx.RepositoryFindAllItem))

			if tt.IsErr && err != nil {
				assert.Equal(t, tt.Err.(*dErr.Error).Type, err.(*dErr.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Len(t, tt.Expected.([]xxx.Props), len(result))
		})
	}
}

func TestXxxGatewayUpdate(t *testing.T) {
	tests := make([]test.Case, 0)

	test1Xxx := test.NewXxx("test1", 1)
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			gateway.Save(context.Background(), &xxx.RepositorySaveItem{Xxx: test1Xxx})
		},
		Ctx: context.Background(),
		Args: &xxx.RepositoryUpdateItem{
			Xxx: test1Xxx,
		},
		Expected: test1Xxx,
		IsErr:    false,
	}
	tests = append(tests, test1)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Setup()

			result, err := gateway.Update(tt.Ctx, tt.Args.(*xxx.RepositoryUpdateItem))

			if tt.IsErr && err != nil {
				assert.Equal(t, tt.Err.(*dErr.Error).Type, err.(*dErr.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.Expected.(xxx.Props), result)
		})
	}
}

func TestXxxGatewayDelete(t *testing.T) {
	tests := make([]test.Case, 0)

	test1Xxx := test.NewXxx("test1", 1)
	test1 := test.Case{
		Name: "正常動作確認",
		Setup: func() {
			gateway.Save(context.Background(), &xxx.RepositorySaveItem{Xxx: test1Xxx})
		},
		Ctx: context.Background(),
		Args: &xxx.RepositoryDeleteItem{
			Id: test1Xxx.Id(),
		},
		Expected: nil,
		IsErr:    false,
	}
	tests = append(tests, test1)

	test2 := test.Case{
		Name:  "存在しないデータを削除",
		Setup: func() {},
		Ctx:   context.Background(),
		Args: &xxx.RepositoryDeleteItem{
			Id: test1Xxx.Id(),
		},
		Expected: nil,
		IsErr:    true,
		Err:      dErr.NewInternalServerError(errors.New(""), ""),
	}
	tests = append(tests, test2)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Setup()

			err := gateway.Delete(tt.Ctx, tt.Args.(*xxx.RepositoryDeleteItem))

			if tt.IsErr && err != nil {
				assert.Equal(t, tt.Err.(*dErr.Error).Type, err.(*dErr.Error).Type)
				return
			} else {
				assert.NoError(t, err)
			}

			_, err = gateway.Find(tt.Ctx, &xxx.RepositoryFindItem{Id: tt.Args.(*xxx.RepositoryDeleteItem).Id})

			if err == nil {
				t.Fail()
			}

		})
	}
}
