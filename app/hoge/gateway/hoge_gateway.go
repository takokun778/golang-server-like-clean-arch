package gateway

import (
	"clean/app/common/gateway"
	c "clean/app/domain/common"
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

func (g *hogeGateway) Find(ctx context.Context, id c.Id) (*hoge.Hoge, error) {
	res, err := g.db.Hoge.Get(ctx, id.Value())
	if err != nil {
		return nil, err
	}

	return hoge.New(
		c.NewId(res.ID),
		hoge.NewName(res.Name),
		hoge.NewNumber(res.Number),
		c.NewTime(res.CreatedAt),
		c.NewTime(res.UpdatedAt),
	), nil
}

func (g *hogeGateway) FindAll(ctx context.Context) (*hoge.HogeList, error) {
	res, err := g.db.Hoge.
		Query().All(ctx)

	if err != nil {
		return nil, err
	}

	list := make([]*hoge.Hoge, 0)

	for _, r := range res {
		list = append(list, hoge.New(
			c.NewId(r.ID),
			hoge.NewName(r.Name),
			hoge.NewNumber(r.Number),
			c.NewTime(r.CreatedAt),
			c.NewTime(r.UpdatedAt),
		))
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

func (g *hogeGateway) Delete(ctx context.Context, id c.Id) error {
	err := g.db.Hoge.DeleteOneID(id.Value()).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
