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

type NewProps struct {
	Id        c.Id
	Name      Name
	Number    Number
	CreatedAt c.Time
	UpdatedAt c.Time
}

func New(props NewProps) *Fuga {
	hoge := new(Fuga)
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

func CreateNew(props CreateNewProps) *Fuga {
	hoge := new(Fuga)
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

func (f *Fuga) Update(props UpdateProps) {
	f.name = props.Name
	f.number = props.Number
	f.updatedAt = c.Now()
}
