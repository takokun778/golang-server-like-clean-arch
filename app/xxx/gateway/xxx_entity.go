package gateway

import (
	"time"

	"xxx/app/domain/xxx"

	"github.com/google/uuid"
)

type XxxEntity struct {
	Id        uuid.UUID
	Name      string
	Number    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Entity(props xxx.Props) XxxEntity {
	return XxxEntity{
		Id:        props.Id().Value(),
		Name:      props.Name().Value(),
		Number:    props.Number().Value(),
		CreatedAt: props.CreatedAt().Value(),
		UpdatedAt: props.UpdatedAt().Value(),
	}
}

func (e XxxEntity) ToProps() xxx.Props {
	id, _ := xxx.ParseId(e.Id.String())
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
