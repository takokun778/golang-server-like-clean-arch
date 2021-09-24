package hoge_test

import (
	"clean/app/domain/common"
	"clean/app/domain/hoge"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	id := common.CreateRandomId()
	name := hoge.NewName("name")
	number := hoge.NewNumber(1)
	now := common.Now()
	createdAt := now
	updatedAt := now

	props := hoge.NewProps{
		Id:        id,
		Name:      name,
		Number:    number,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	result := hoge.New(props)

	assert.Equal(t, result.Id(), id)
	assert.Equal(t, result.Name(), name)
	assert.Equal(t, result.Number(), number)
	assert.Equal(t, result.CreatedAt(), createdAt)
	assert.Equal(t, result.UpdatedAt(), updatedAt)
}

func TestNewCreate(t *testing.T) {
	name := hoge.NewName("create")
	number := hoge.NewNumber(1)

	props := hoge.CreateNewProps{
		Name:   name,
		Number: number,
	}

	result := hoge.CreateNew(props)

	assert.Equal(t, result.Name(), name)
	assert.Equal(t, result.Number(), number)
	assert.Equal(t, result.CreatedAt(), result.UpdatedAt())
}

func TestUpdate(t *testing.T) {
	createNewProps := hoge.CreateNewProps{
		Name:   hoge.NewName("new"),
		Number: hoge.NewNumber(1),
	}

	result := hoge.CreateNew(createNewProps)

	time.Sleep(time.Millisecond)

	name := hoge.NewName("update")
	number := hoge.NewNumber(2)

	props := hoge.UpdateProps{
		Name:   name,
		Number: number,
	}

	result.Update(props)

	assert.Equal(t, result.Name(), name)
	assert.Equal(t, result.Number(), number)
	assert.Greater(t, result.UpdatedAt().Value().UnixNano(), result.CreatedAt().Value().UnixNano())
}
