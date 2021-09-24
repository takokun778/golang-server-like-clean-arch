package hoge

import (
	c "clean/app/domain/common"
)

type Hoge struct {
	id        c.Id
	name      Name
	number    Number
	createdAt c.Time
	updatedAt c.Time
}

func (h *Hoge) Id() c.Id {
	return h.id
}

func (h *Hoge) Name() Name {
	return h.name
}

func (h *Hoge) Number() Number {
	return h.number
}

func (h *Hoge) CreatedAt() c.Time {
	return h.createdAt
}

func (h *Hoge) UpdatedAt() c.Time {
	return h.updatedAt
}

func New(
	id c.Id,
	name Name,
	number Number,
	createdAt c.Time,
	updatedAt c.Time,
) *Hoge {
	hoge := new(Hoge)
	hoge.id = id
	hoge.name = name
	hoge.number = number
	hoge.createdAt = createdAt
	hoge.updatedAt = updatedAt
	return hoge
}

func CreateNew(name Name, number Number) *Hoge {
	now := c.Now()
	return New(
		c.CreateRandomId(),
		name,
		number,
		now,
		now,
	)
}

func (h *Hoge) Update(name Name, number Number) {
	h.name = name
	h.number = number
	h.updatedAt = c.Now()
}
