package fuga

import (
	"clean/app/domain/common"
)

type Fuga struct {
	id        common.Id
	name      Name
	number    Number
	createdAt common.Time
	updatedAt common.Time
}

func (f *Fuga) Id() common.Id {
	return f.id
}

func (f *Fuga) Name() Name {
	return f.name
}

func (f *Fuga) Number() Number {
	return f.number
}

func (f *Fuga) CreatedAt() common.Time {
	return f.createdAt
}

func (f *Fuga) UpdatedAt() common.Time {
	return f.updatedAt
}

func New(
	id common.Id,
	name Name,
	number Number,
	createdAt common.Time,
	updatedAt common.Time,
) *Fuga {
	hoge := new(Fuga)
	hoge.id = id
	hoge.name = name
	hoge.number = number
	hoge.createdAt = createdAt
	hoge.updatedAt = updatedAt
	return hoge
}

func CreateNew(name Name, number Number) *Fuga {
	now := common.Now()
	return New(
		common.CreateRandomId(),
		name,
		number,
		now,
		now,
	)
}

func (f *Fuga) Update(name Name, number Number) *Fuga {
	return New(
		f.id,
		name,
		number,
		f.createdAt,
		common.Now(),
	)
}
