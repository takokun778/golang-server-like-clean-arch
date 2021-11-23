package xxx

type Props struct {
	props
}

type props struct {
	id
	name
	number
	createdAt time
	updatedAt time
}

func (p props) Id() id {
	return p.id
}

func (p props) Name() name {
	return p.name
}

func (p props) Number() number {
	return p.number
}

func (p props) CreatedAt() time {
	return p.createdAt
}

func (p props) UpdatedAt() time {
	return p.updatedAt
}

func NewProps(
	id id,
	name name,
	number number,
	createdAt time,
	updatedAt time,
) Props {
	props := new(Props)
	props.id = id
	props.name = name
	props.number = number
	props.createdAt = createdAt
	props.updatedAt = updatedAt
	return *props
}
