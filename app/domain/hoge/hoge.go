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

type NewProps struct {
	Id        c.Id
	Name      Name
	Number    Number
	CreatedAt c.Time
	UpdatedAt c.Time
}

func New(props NewProps) *Hoge {
	hoge := new(Hoge)
	hoge.id = props.Id
	hoge.name = props.Name
	hoge.number = props.Number
	hoge.createdAt = props.CreatedAt
	hoge.updatedAt = props.UpdatedAt
	return hoge
}

type CreateNewProps struct {
	Name   Name
	Number Number
}

func CreateNew(props CreateNewProps) *Hoge {
	hoge := new(Hoge)
	hoge.id = c.CreateRandomId()
	hoge.name = props.Name
	hoge.number = props.Number
	hoge.createdAt = c.Now()
	hoge.updatedAt = c.Now()
	return hoge
}

type UpdateProps struct {
	Name   Name
	Number Number
}

func (h *Hoge) Update(props UpdateProps) {
	h.name = props.Name
	h.number = props.Number
	h.updatedAt = c.Now()
}
