package controller_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	cg "clean/app/common/gateway"

	fc "clean/app/fuga/controller"
	fg "clean/app/fuga/gateway"
	fu "clean/app/fuga/usecase"
	pbFuga "clean/proto/fuga"

	"github.com/stretchr/testify/assert"
)

var ctx = context.TODO()
var gateway = fg.NewFugaGateway()
var usecase = fu.NewFugaUsecase(gateway)
var controller = fc.NewFugaController(usecase)

func TestMain(m *testing.M) {
	// テスト前処理
	createTable()

	// テスト実行
	m.Run()
}

func createTable() {
	database := cg.DatabaseConnect()

	ctx := context.TODO()

	if err := database.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("create table completed")
}

func resetTable() {
	datasource := fmt.Sprintf("host=%s port=%s user=%s dbname= %s password=%s sslmode=disable", "localhost", "54321", "test", "test", "test")

	database, err := sql.Open("postgres", datasource)

	if err != nil {
		log.Fatal(err)
	}

	database.Exec("DELETE FROM fugas")
}

func TestCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := new(pbFuga.CreateRequest)
		req.Name = "create"
		req.Number = 1

		res, err := controller.Create(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, res.Fuga.Name, "create")
		assert.Equal(t, res.Fuga.Number, req.Number)

	})
	resetTable()
}

func TestFetch(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dataReq := new(pbFuga.CreateRequest)
		dataReq.Name = "fetch"
		dataReq.Number = 1

		dataRes, err := controller.Create(ctx, dataReq)
		assert.NoError(t, err)

		req := new(pbFuga.FetchRequest)
		req.Id = dataRes.Fuga.Id

		res, err := controller.Fetch(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, res.Fuga.Id, req.Id)
		assert.Equal(t, res.Fuga.Name, dataReq.Name)
		assert.Equal(t, res.Fuga.Number, dataReq.Number)
		assert.Equal(t, res.Fuga.CreatedAt, dataRes.Fuga.CreatedAt)
		assert.Equal(t, res.Fuga.UpdatedAt, dataRes.Fuga.UpdatedAt)
	})
	resetTable()
}

func TestFetchAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dataReq1 := new(pbFuga.CreateRequest)
		dataReq1.Name = "all"
		dataReq1.Number = 1

		_, err := controller.Create(ctx, dataReq1)
		assert.NoError(t, err)

		dataReq2 := new(pbFuga.CreateRequest)
		dataReq2.Name = "all"
		dataReq2.Number = 1

		_, err = controller.Create(ctx, dataReq2)
		assert.NoError(t, err)

		req := new(pbFuga.FetchAllRequest)

		res, err := controller.FetchAll(ctx, req)

		assert.NoError(t, err)
		assert.Len(t, res.Fugas, 2)
	})
	resetTable()
}

func TestUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dataReq := new(pbFuga.CreateRequest)
		dataReq.Name = "src"
		dataReq.Number = 1

		dataRes, err := controller.Create(ctx, dataReq)
		assert.NoError(t, err)

		req := new(pbFuga.UpdateRequest)
		req.Id = dataRes.Fuga.Id
		req.Name = "dst"
		req.Number = 2

		res, err := controller.Update(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, res.Fuga.Id, req.Id)
		assert.Equal(t, res.Fuga.Name, req.Name)
		assert.Equal(t, res.Fuga.Number, req.Number)
		assert.Equal(t, res.Fuga.CreatedAt, dataRes.Fuga.CreatedAt)
		assert.Greater(t, res.Fuga.UpdatedAt.Nanos, dataRes.Fuga.UpdatedAt.Nanos)

	})
	resetTable()
}

func TestDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dataReq := new(pbFuga.CreateRequest)
		dataReq.Name = "delete"
		dataReq.Number = 11

		dataRes, err := controller.Create(ctx, dataReq)
		assert.NoError(t, err)

		req := new(pbFuga.DeleteRequest)
		req.Id = dataRes.Fuga.Id

		res, err := controller.Delete(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, res.Fuga.Id, req.Id)

		fetchReq := new(pbFuga.FetchRequest)
		fetchReq.Id = res.Fuga.Id

		fetchRes, err := controller.Fetch(ctx, fetchReq)
		assert.Nil(t, fetchRes)
		assert.Error(t, err)
	})
	resetTable()
}
