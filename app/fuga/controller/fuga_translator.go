package controller

import (
	"clean/app/domain/fuga"
	pbFuga "clean/proto/fuga"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type FugaTranslator struct{}

func (FugaTranslator) ToProto(fuga *fuga.Fuga) *pbFuga.Fuga {
	proto := new(pbFuga.Fuga)
	proto.Id = fuga.Id().Value().String()
	proto.Name = fuga.Name().Value()
	proto.Number = fuga.Number().Value()
	proto.CreatedAt = timestamppb.New(fuga.CreatedAt().Value())
	proto.UpdatedAt = timestamppb.New(fuga.UpdatedAt().Value())
	return proto
}

func (FugaTranslator) ToProtoList(fugaList *fuga.FugaList) []*pbFuga.Fuga {
	proto := make([]*pbFuga.Fuga, 0)

	for _, fuga := range *fugaList {
		proto = append(proto, FugaTranslator{}.ToProto(fuga))
	}

	return proto
}
