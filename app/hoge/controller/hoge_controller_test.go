package controller_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	cg "clean/app/common/gateway"

	hc "clean/app/hoge/controller"
	hg "clean/app/hoge/gateway"
	hu "clean/app/hoge/usecase"
	pbHoge "clean/proto/hoge"

	"github.com/stretchr/testify/assert"
)

var ctx = context.TODO()
var gateway = hg.NewHogeGateway()
var usecase = hu.NewHogeUsecase(gateway)
var controller = hc.NewHogeController(usecase)

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

	database.Exec("DELETE FROM hoges")
}

func TestCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := new(pbHoge.CreateRequest)
		req.Name = "create"
		req.Number = 1

		res, err := controller.Create(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, res.Hoge.Name, "create")
		assert.Equal(t, res.Hoge.Number, req.Number)

	})
	resetTable()
}

func TestFetch(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dataReq := new(pbHoge.CreateRequest)
		dataReq.Name = "fetch"
		dataReq.Number = 1

		dataRes, err := controller.Create(ctx, dataReq)
		assert.NoError(t, err)

		req := new(pbHoge.FetchRequest)
		req.Id = dataRes.Hoge.Id

		res, err := controller.Fetch(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, res.Hoge.Id, req.Id)
		assert.Equal(t, res.Hoge.Name, dataReq.Name)
		assert.Equal(t, res.Hoge.Number, dataReq.Number)
		assert.Equal(t, res.Hoge.CreatedAt, dataRes.Hoge.CreatedAt)
		assert.Equal(t, res.Hoge.UpdatedAt, dataRes.Hoge.UpdatedAt)
	})
	resetTable()
}

func TestFetchAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dataReq1 := new(pbHoge.CreateRequest)
		dataReq1.Name = "all"
		dataReq1.Number = 1

		_, err := controller.Create(ctx, dataReq1)
		assert.NoError(t, err)

		dataReq2 := new(pbHoge.CreateRequest)
		dataReq2.Name = "all"
		dataReq2.Number = 1

		_, err = controller.Create(ctx, dataReq2)
		assert.NoError(t, err)

		req := new(pbHoge.FetchAllRequest)

		res, err := controller.FetchAll(ctx, req)

		assert.NoError(t, err)
		assert.Len(t, res.Hoges, 2)
	})
	resetTable()
}

func TestUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dataReq := new(pbHoge.CreateRequest)
		dataReq.Name = "src"
		dataReq.Number = 1

		dataRes, err := controller.Create(ctx, dataReq)
		assert.NoError(t, err)

		req := new(pbHoge.UpdateRequest)
		req.Id = dataRes.Hoge.Id
		req.Name = "dst"
		req.Number = 2

		res, err := controller.Update(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, res.Hoge.Id, req.Id)
		assert.Equal(t, res.Hoge.Name, req.Name)
		assert.Equal(t, res.Hoge.Number, req.Number)
		assert.Equal(t, res.Hoge.CreatedAt, dataRes.Hoge.CreatedAt)
		assert.Greater(t, res.Hoge.UpdatedAt.Nanos, dataRes.Hoge.UpdatedAt.Nanos)

	})
	resetTable()
}

func TestDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dataReq := new(pbHoge.CreateRequest)
		dataReq.Name = "delete"
		dataReq.Number = 11

		dataRes, err := controller.Create(ctx, dataReq)
		assert.NoError(t, err)

		req := new(pbHoge.DeleteRequest)
		req.Id = dataRes.Hoge.Id

		res, err := controller.Delete(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, res.Hoge.Id, req.Id)

		fetchReq := new(pbHoge.FetchRequest)
		fetchReq.Id = res.Hoge.Id

		fetchRes, err := controller.Fetch(ctx, fetchReq)
		assert.Nil(t, fetchRes)
		assert.Error(t, err)
	})
	resetTable()
}
