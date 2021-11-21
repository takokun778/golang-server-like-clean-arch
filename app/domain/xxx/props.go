package xxx

import (
	"xxx/app/domain/common"
)

type Props struct {
	props
}

type props struct {
	id        common.Id
	name      Name
	number    Number
	createdAt common.Time
	updatedAt common.Time
}

func (p props) Id() common.Id {
	return p.id
}

func (p props) Name() Name {
	return p.name
}

func (p props) Number() Number {
	return p.number
}

func (p props) CreatedAt() common.Time {
	return p.createdAt
}

func (p props) UpdatedAt() common.Time {
	return p.updatedAt
}

func NewProps(
	id common.Id,
	name Name,
	number Number,
	createdAt common.Time,
	updatedAt common.Time,
) Props {
	props := new(Props)
	props.id = id
	props.name = name
	props.number = number
	props.createdAt = createdAt
	props.updatedAt = updatedAt
	return *props
}
