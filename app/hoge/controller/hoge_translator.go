package controller

import (
	"clean/app/domain/hoge"
	pbHoge "clean/proto/hoge"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type HogeTranslator struct{}

func (HogeTranslator) ToProto(hoge *hoge.Hoge) *pbHoge.Hoge {
	proto := new(pbHoge.Hoge)
	proto.Id = hoge.Id().Value().String()
	proto.Name = hoge.Name().Value()
	proto.Number = hoge.Number().Value()
	proto.CreatedAt = timestamppb.New(hoge.CreatedAt().Value())
	proto.UpdatedAt = timestamppb.New(hoge.UpdatedAt().Value())
	return proto
}

func (HogeTranslator) ToProtoList(hogeList *hoge.HogeList) []*pbHoge.Hoge {
	proto := make([]*pbHoge.Hoge, 0)

	for _, hoge := range hogeList.Items() {
		proto = append(proto, HogeTranslator{}.ToProto(hoge))
	}

	return proto
}
