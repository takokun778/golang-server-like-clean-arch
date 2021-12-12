package xxx_test

import (
	"context"
	"testing"

	"xxx/app/infra/ent"
)

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
