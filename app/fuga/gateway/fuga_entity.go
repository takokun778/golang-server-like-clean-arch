package gateway

import (
	"clean/app/domain/common"
	"clean/app/domain/fuga"
	"clean/ent"
)

type FugaEntity ent.Fuga

func (e *FugaEntity) ToValues() fuga.Values {
	return fuga.NewValues(
		common.Id(e.ID),
		fuga.Name(e.Name),
		fuga.Number(e.Number),
		common.Time(e.CreatedAt),
		common.Time(e.UpdatedAt),
	)
}
