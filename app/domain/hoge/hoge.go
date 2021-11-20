package hoge

import (
	"clean/app/domain/common"
)

type Hoge struct {
	values
}

func (h Hoge) Values() Values {
	return Values{h.values}
}

func constructor(
	id common.Id,
	name Name,
	number Number,
	createdAt common.Time,
	updatedAt common.Time,
) Hoge {
	hoge := new(Hoge)
	hoge.id = id
	hoge.name = name
	hoge.number = number
	hoge.createdAt = createdAt
	hoge.updatedAt = updatedAt
	return *hoge
}

func Reconstruct(values Values) Hoge {
	return constructor(
		values.id,
		values.name,
		values.number,
		values.createdAt,
		values.updatedAt,
	)
}

func Create(name Name, number Number) Hoge {
	now := common.Now()
	return constructor(
		common.CreateRandomId(),
		name,
		number,
		now,
		now,
	)
}

func (h Hoge) Update(name Name, number Number) Hoge {
	return constructor(
		h.id,
		name,
		number,
		h.createdAt,
		common.Now(),
	)
}
