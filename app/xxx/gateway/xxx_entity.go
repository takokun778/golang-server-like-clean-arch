package gateway

import (
	"xxx/app/domain/common"
	"xxx/app/domain/xxx"
	"xxx/ent"
)

type XxxEntity ent.Xxx

func (e XxxEntity) ToProps() xxx.Props {
	return xxx.NewProps(
		common.Id(e.ID),
		xxx.Name(e.Name),
		xxx.Number(e.Number),
		common.Time(e.CreatedAt),
		common.Time(e.UpdatedAt),
	)
}
