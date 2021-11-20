package xxx

import (
	"xxx/app/domain/common"
)

type Values struct {
	values
}

type values struct {
	id        common.Id
	name      Name
	number    Number
	createdAt common.Time
	updatedAt common.Time
}

func (v values) Id() common.Id {
	return v.id
}

func (v values) Name() Name {
	return v.name
}

func (v values) Number() Number {
	return v.number
}

func (v values) CreatedAt() common.Time {
	return v.createdAt
}

func (v values) UpdatedAt() common.Time {
	return v.updatedAt
}

func NewValues(
	id common.Id,
	name Name,
	number Number,
	createdAt common.Time,
	updatedAt common.Time,
) Values {
	values := new(Values)
	values.id = id
	values.name = name
	values.number = number
	values.createdAt = createdAt
	values.updatedAt = updatedAt
	return *values
}
