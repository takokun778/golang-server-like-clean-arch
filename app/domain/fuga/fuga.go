package fuga

import (
	c "clean/app/domain/common"
)

type Fuga struct {
	id        c.Id
	name      Name
	number    Number
	createdAt c.Time
	updatedAt c.Time
}

func (f *Fuga) Id() c.Id {
	return f.id
}

func (f *Fuga) Name() Name {
	return f.name
}

func (f *Fuga) Number() Number {
	return f.number
}

func (f *Fuga) CreatedAt() c.Time {
	return f.createdAt
}

func (f *Fuga) UpdatedAt() c.Time {
	return f.updatedAt
}

func New(
	id c.Id,
	name Name,
	number Number,
	createdAt c.Time,
	updatedAt c.Time,
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
	now := c.Now()
	return New(
		c.CreateRandomId(),
		name,
		number,
		now,
		now,
	)
}

func (f *Fuga) Update(name Name, number Number) {
	f.name = name
	f.number = number
	f.updatedAt = c.Now()
}
