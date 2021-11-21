package xxx

import (
	"xxx/app/domain/common"
)

type Xxx struct {
	props
}

func (x Xxx) Props() Props {
	return Props(x)
}

func constructor(
	id common.Id,
	name Name,
	number Number,
	createdAt common.Time,
	updatedAt common.Time,
) Xxx {
	xxx := new(Xxx)
	xxx.id = id
	xxx.name = name
	xxx.number = number
	xxx.createdAt = createdAt
	xxx.updatedAt = updatedAt
	return *xxx
}

func Reconstruct(values Props) Xxx {
	return constructor(
		values.id,
		values.name,
		values.number,
		values.createdAt,
		values.updatedAt,
	)
}

func Create(name Name, number Number) Xxx {
	now := common.Now()
	return constructor(
		common.CreateRandomId(),
		name,
		number,
		now,
		now,
	)
}

func (x Xxx) Update(name Name, number Number) Xxx {
	return constructor(
		x.id,
		name,
		number,
		x.createdAt,
		common.Now(),
	)
}
