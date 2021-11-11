package gateway

import (
	"clean/app/common/gateway"
	"clean/app/domain/common"
	"clean/app/domain/hoge"
	"context"
)

type hogeGateway struct {
	db *gateway.Database
}

func NewHogeGateway() hoge.Repository {
	db := gateway.DatabaseConnect()
	gw := new(hogeGateway)
	gw.db = db
	return gw
}

func (g *hogeGateway) Save(ctx context.Context, hoge *hoge.Hoge) (*hoge.Hoge, error) {
	_, err := g.db.Hoge.
		Create().
		SetID(hoge.Id().Value()).
		SetName(hoge.Name().Value()).
		SetNumber(hoge.Number().Value()).
		SetCreatedAt(hoge.CreatedAt().Value()).
		SetUpdatedAt(hoge.UpdatedAt().Value()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return hoge, nil
}

func (g *hogeGateway) Find(ctx context.Context, id common.Id) (*hoge.Hoge, error) {
	res, err := g.db.Hoge.Get(ctx, id.Value())
	if err != nil {
		return nil, err
	}

	entity := HogeEntity(*res)

	return entity.ToDomain(), nil
}

func (g *hogeGateway) FindAll(ctx context.Context) (*hoge.HogeList, error) {
	res, err := g.db.Hoge.
		Query().All(ctx)

	if err != nil {
		return nil, err
	}

	list := make([]*hoge.Hoge, 0)

	for _, r := range res {
		entity := HogeEntity(*r)
		list = append(list, entity.ToDomain())
	}

	return hoge.NewList(list), nil
}

func (g *hogeGateway) Update(ctx context.Context, hoge *hoge.Hoge) (*hoge.Hoge, error) {
	_, err := g.db.Hoge.
		Update().
		SetName(hoge.Name().Value()).
		SetNumber(hoge.Number().Value()).
		SetUpdatedAt(hoge.UpdatedAt().Value()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return hoge, nil
}

func (g *hogeGateway) Delete(ctx context.Context, id common.Id) error {
	err := g.db.Hoge.DeleteOneID(id.Value()).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
