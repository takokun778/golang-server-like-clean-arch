package hoge

import (
	"clean/app/domain/common"
)

type Hoge struct {
	id        common.Id
	name      Name
	number    Number
	createdAt common.Time
	updatedAt common.Time
}

func (h *Hoge) Id() common.Id {
	return h.id
}

func (h *Hoge) Name() Name {
	return h.name
}

func (h *Hoge) Number() Number {
	return h.number
}

func (h *Hoge) CreatedAt() common.Time {
	return h.createdAt
}

func (h *Hoge) UpdatedAt() common.Time {
	return h.updatedAt
}

func New(
	id common.Id,
	name Name,
	number Number,
	createdAt common.Time,
	updatedAt common.Time,
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
	now := common.Now()
	return New(
		common.CreateRandomId(),
		name,
		number,
		now,
		now,
	)
}

func (h *Hoge) Update(name Name, number Number) {
	h.name = name
	h.number = number
	h.updatedAt = common.Now()
}
