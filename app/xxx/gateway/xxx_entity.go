package gateway

import (
	"xxx/app/domain/xxx"
	"xxx/ent"
)

type XxxEntity ent.Xxx

func Entity(props xxx.Props) XxxEntity {
	return XxxEntity{
		ID:        props.Id().Value(),
		Name:      props.Name().Value(),
		Number:    props.Number().Value(),
		CreatedAt: props.CreatedAt().Value(),
		UpdatedAt: props.UpdatedAt().Value(),
	}
}

func (e XxxEntity) ToProps() xxx.Props {
	id, _ := xxx.ParseId(e.ID.String())
	name, _ := xxx.NewName(e.Name)
	number, _ := xxx.NewNumber(e.Number)
	return xxx.NewProps(
		id,
		name,
		number,
		xxx.NewTime(e.CreatedAt),
		xxx.NewTime(e.UpdatedAt),
	)
}
