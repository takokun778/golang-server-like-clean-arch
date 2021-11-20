package gateway

import (
	"clean/app/domain/common"
	"clean/app/domain/hoge"
	"clean/ent"
)

type HogeEntity ent.Hoge

func (e *HogeEntity) ToValues() hoge.Values {
	return hoge.NewValues(
		common.Id(e.ID),
		hoge.Name(e.Name),
		hoge.Number(e.Number),
		common.Time(e.CreatedAt),
		common.Time(e.UpdatedAt),
	)
}
