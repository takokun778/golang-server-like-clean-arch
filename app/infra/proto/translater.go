package proto

import (
	"xxx/app/domain/xxx"
	pbXxx "xxx/proto/xxx"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (*Proto) Translate(xxx xxx.Props) *pbXxx.Xxx {
	proto := new(pbXxx.Xxx)
	proto.Id = xxx.Id().Value().String()
	proto.Name = xxx.Name().Value()
	proto.Number = int32(xxx.Number().Value())
	proto.CreatedAt = timestamppb.New(xxx.CreatedAt().Value())
	proto.UpdatedAt = timestamppb.New(xxx.UpdatedAt().Value())
	return proto
}

func (p *Proto) TranslateArray(xxxs []xxx.Props) []*pbXxx.Xxx {
	proto := make([]*pbXxx.Xxx, 0)

	for _, xxx := range xxxs {
		proto = append(proto, p.Translate(xxx))
	}

	return proto
}
