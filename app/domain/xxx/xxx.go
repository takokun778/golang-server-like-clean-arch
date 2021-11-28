package xxx

type xxx struct {
	props
}

func (x xxx) Props() Props {
	return Props(x)
}

func constructor(
	id id,
	name name,
	number number,
	createdAt time,
	updatedAt time,
) xxx {
	xxx := new(xxx)
	xxx.id = id
	xxx.name = name
	xxx.number = number
	xxx.createdAt = createdAt
	xxx.updatedAt = updatedAt
	return *xxx
}

func Reconstruct(values Props) xxx {
	return constructor(
		values.id,
		values.name,
		values.number,
		values.createdAt,
		values.updatedAt,
	)
}

func Create(name name, number number) xxx {
	now := TimeNow()
	return constructor(
		CreateRandomId(),
		name,
		number,
		now,
		now,
	)
}

func (x xxx) Update(name name, number number) xxx {
	return constructor(
		x.id,
		name,
		number,
		x.createdAt,
		TimeNow(),
	)
}
