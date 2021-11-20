package gateway

import (
	"xxx/app/domain/common"
	"xxx/app/domain/xxx"
	"xxx/ent"
)

type XxxEntity ent.Xxx

func (e XxxEntity) ToValues() xxx.Values {
	return xxx.NewValues(
		common.Id(e.ID),
		xxx.Name(e.Name),
		xxx.Number(e.Number),
		common.Time(e.CreatedAt),
		common.Time(e.UpdatedAt),
	)
}
