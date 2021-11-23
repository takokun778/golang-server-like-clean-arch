package xxx

type Xxx struct {
	props
}

func (x Xxx) Props() Props {
	return Props(x)
}

func constructor(
	id id,
	name name,
	number number,
	createdAt time,
	updatedAt time,
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

func Create(name name, number number) Xxx {
	now := TimeNow()
	return constructor(
		CreateRandomId(),
		name,
		number,
		now,
		now,
	)
}

func (x Xxx) Update(name name, number number) Xxx {
	return constructor(
		x.id,
		name,
		number,
		x.createdAt,
		TimeNow(),
	)
}
