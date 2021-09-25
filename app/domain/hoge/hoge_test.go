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
	name := hoge.Name("name")
	number := hoge.Number(1)
	now := common.Now()
	createdAt := now
	updatedAt := now

	result := hoge.New(
		id,
		name,
		number,
		now,
		now,
	)

	assert.Equal(t, result.Id(), id)
	assert.Equal(t, result.Name(), name)
	assert.Equal(t, result.Number(), number)
	assert.Equal(t, result.CreatedAt(), createdAt)
	assert.Equal(t, result.UpdatedAt(), updatedAt)
}

func TestNewCreate(t *testing.T) {
	name := hoge.Name("create")
	number := hoge.Number(1)

	result := hoge.CreateNew(name, number)

	assert.Equal(t, result.Name(), name)
	assert.Equal(t, result.Number(), number)
	assert.Equal(t, result.CreatedAt(), result.UpdatedAt())
}

func TestUpdate(t *testing.T) {
	result := hoge.CreateNew(hoge.Name("new"), hoge.Number(1))

	time.Sleep(time.Millisecond)

	name := hoge.Name("update")
	number := hoge.Number(2)

	result.Update(name, number)

	assert.Equal(t, result.Name(), name)
	assert.Equal(t, result.Number(), number)
	assert.Greater(t, result.UpdatedAt().Value().UnixNano(), result.CreatedAt().Value().UnixNano())
}
