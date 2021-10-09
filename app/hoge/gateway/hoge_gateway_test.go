package gateway_test

import (
	"context"
	"log"
	"testing"

	cg "clean/app/common/gateway"
	dh "clean/app/domain/hoge"
	hg "clean/app/hoge/gateway"

	"github.com/stretchr/testify/assert"
)

var ctx = context.Background()
var gateway = hg.NewHogeGateway()

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

	database.Fuga.Delete().Exec(ctx)
}

func TestHogeGatewaySave(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		fuga := dh.CreateNew(
			dh.Name("fuga"),
			dh.Number(1),
		)
		result, err := gateway.Save(ctx, fuga)
		if err != nil {
			log.Fatal(err)
		}
		assert.Equal(t, result, fuga)
	})
}

func TestHogeGatewayFind(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		fuga := dh.CreateNew(
			dh.Name("fuga"),
			dh.Number(1),
		)
		_, err := gateway.Save(ctx, fuga)
		if err != nil {
			log.Fatal(err)
		}
		result, err := gateway.Find(ctx, fuga.Id())
		if err != nil {
			log.Fatal(err)
		}
		assert.Equal(t, result, fuga)
	})
}

func TestHogeGatewayUpdate(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		fuga := dh.CreateNew(
			dh.Name("fuga"),
			dh.Number(1),
		)
		_, err := gateway.Save(ctx, fuga)
		if err != nil {
			log.Fatal(err)
		}
		fuga.Update(
			dh.Name("fugafuga"),
			dh.Number(2),
		)
		result, err := gateway.Update(ctx, fuga)
		if err != nil {
			log.Fatal(err)
		}
		assert.Equal(t, result, fuga)
		foundFuga, err := gateway.Find(ctx, fuga.Id())
		if err != nil {
			log.Fatal(err)
		}
		assert.Equal(t, foundFuga, fuga)
	})
}

func TestHogeGatewayDelete(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		fuga := dh.CreateNew(
			dh.Name("fuga"),
			dh.Number(1),
		)
		_, err := gateway.Save(ctx, fuga)
		if err != nil {
			log.Fatal(err)
		}
		delErr := gateway.Delete(ctx, fuga.Id())
		if delErr != nil {
			log.Fatal(err)
		}
		_, finErr := gateway.Find(ctx, fuga.Id())
		if finErr == nil {
			log.Fatal(err)
		}
	})
}
