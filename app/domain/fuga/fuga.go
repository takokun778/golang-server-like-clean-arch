package fuga

import (
	"clean/app/domain/common"
)

type Fuga struct {
	values
}

func (f Fuga) Values() Values {
	return Values{f.values}
}

func constructor(
	id common.Id,
	name Name,
	number Number,
	createdAt common.Time,
	updatedAt common.Time,
) Fuga {
	fuga := new(Fuga)
	fuga.id = id
	fuga.name = name
	fuga.number = number
	fuga.createdAt = createdAt
	fuga.updatedAt = updatedAt
	return *fuga
}

func Reconstruct(values Values) Fuga {
	return constructor(
		values.id,
		values.name,
		values.number,
		values.createdAt,
		values.updatedAt,
	)
}

func Create(name Name, number Number) Fuga {
	now := common.Now()
	return constructor(
		common.CreateRandomId(),
		name,
		number,
		now,
		now,
	)
}

func (f Fuga) Update(name Name, number Number) Fuga {
	return constructor(
		f.id,
		name,
		number,
		f.createdAt,
		common.Now(),
	)
}
