package gateway

import (
	"clean/app/domain/common"
	"clean/app/domain/fuga"
	"clean/ent"
)

type FugaEntity ent.Fuga

func (e *FugaEntity) ToDomain() *fuga.Fuga {
	return fuga.New(
		common.Id(e.ID),
		fuga.Name(e.Name),
		fuga.Number(e.Number),
		common.Time(e.CreatedAt),
		common.Time(e.UpdatedAt),
	)
}
