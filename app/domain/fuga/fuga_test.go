package fuga_test

import (
	"clean/app/domain/common"
	"clean/app/domain/fuga"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	id := common.CreateRandomId()
	name := fuga.NewName("name")
	number := fuga.NewNumber(1)
	now := common.Now()
	createdAt := now
	updatedAt := now

	props := fuga.NewProps{
		Id:        id,
		Name:      name,
		Number:    number,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	result := fuga.New(props)

	assert.Equal(t, result.Id(), id)
	assert.Equal(t, result.Name(), name)
	assert.Equal(t, result.Number(), number)
	assert.Equal(t, result.CreatedAt(), createdAt)
	assert.Equal(t, result.UpdatedAt(), updatedAt)
}

func TestNewCreate(t *testing.T) {
	name := fuga.NewName("create")
	number := fuga.NewNumber(1)

	props := fuga.CreateNewProps{
		Name:   name,
		Number: number,
	}

	result := fuga.CreateNew(props)

	assert.Equal(t, result.Name(), name)
	assert.Equal(t, result.Number(), number)
	assert.Equal(t, result.CreatedAt(), result.UpdatedAt())
}

func TestUpdate(t *testing.T) {
	createNewProps := fuga.CreateNewProps{
		Name:   fuga.NewName("new"),
		Number: fuga.NewNumber(1),
	}

	result := fuga.CreateNew(createNewProps)

	time.Sleep(time.Millisecond)

	name := fuga.NewName("update")
	number := fuga.NewNumber(2)

	props := fuga.UpdateProps{
		Name:   name,
		Number: number,
	}

	result.Update(props)

	assert.Equal(t, result.Name(), name)
	assert.Equal(t, result.Number(), number)
	assert.Greater(t, result.UpdatedAt().Value().UnixNano(), result.CreatedAt().Value().UnixNano())
}
