package gateway

import (
	"clean/app/common/gateway"
	c "clean/app/domain/common"
	"clean/app/domain/fuga"
	"context"
)

type fugaGateway struct {
	db *gateway.Database
}

func NewFugaGateway() fuga.Repository {
	db := gateway.DatabaseConnect()
	gw := new(fugaGateway)
	gw.db = db
	return gw
}

func (g *fugaGateway) Save(ctx context.Context, fuga *fuga.Fuga) (*fuga.Fuga, error) {
	_, err := g.db.Fuga.
		Create().
		SetID(fuga.Id().Value()).
		SetName(fuga.Name().Value()).
		SetNumber(fuga.Number().Value()).
		SetCreatedAt(fuga.CreatedAt().Value()).
		SetUpdatedAt(fuga.UpdatedAt().Value()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return fuga, nil
}

func (g *fugaGateway) Find(ctx context.Context, id c.Id) (*fuga.Fuga, error) {
	res, err := g.db.Fuga.Get(ctx, id.Value())
	if err != nil {
		return nil, err
	}

	return fuga.New(
		c.NewId(res.ID),
		fuga.NewName(res.Name),
		fuga.NewNumber(res.Number),
		c.NewTime(res.CreatedAt),
		c.NewTime(res.UpdatedAt),
	), nil
}

func (g *fugaGateway) FindAll(ctx context.Context) (*fuga.FugaList, error) {
	res, err := g.db.Fuga.
		Query().All(ctx)

	if err != nil {
		return nil, err
	}

	list := make([]*fuga.Fuga, 0)

	for _, r := range res {
		list = append(list, fuga.New(
			c.NewId(r.ID),
			fuga.NewName(r.Name),
			fuga.NewNumber(r.Number),
			c.NewTime(r.CreatedAt),
			c.NewTime(r.UpdatedAt),
		))
	}

	return fuga.NewList(list), nil
}

func (g *fugaGateway) Update(ctx context.Context, fuga *fuga.Fuga) (*fuga.Fuga, error) {
	_, err := g.db.Fuga.
		Update().
		SetName(fuga.Name().Value()).
		SetNumber(fuga.Number().Value()).
		SetUpdatedAt(fuga.UpdatedAt().Value()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return fuga, nil
}

func (g *fugaGateway) Delete(ctx context.Context, id c.Id) error {
	err := g.db.Fuga.DeleteOneID(id.Value()).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
