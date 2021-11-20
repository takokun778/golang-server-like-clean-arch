package gateway

import (
	"clean/app/common/gateway"
	"clean/app/domain/common"
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

func (g *fugaGateway) Save(ctx context.Context, values fuga.Values) (fuga.Values, error) {
	_, err := g.db.Fuga.
		Create().
		SetID(values.Id().Value()).
		SetName(values.Name().Value()).
		SetNumber(values.Number().Value()).
		SetCreatedAt(values.CreatedAt().Value()).
		SetUpdatedAt(values.UpdatedAt().Value()).
		Save(ctx)

	if err != nil {
		return fuga.Values{}, err
	}

	return values, nil
}

func (g *fugaGateway) Find(ctx context.Context, id common.Id) (fuga.Values, error) {
	res, err := g.db.Fuga.Get(ctx, id.Value())

	if err != nil {
		return fuga.Values{}, err
	}

	entity := FugaEntity(*res)

	return entity.ToValues(), nil
}

func (g *fugaGateway) FindAll(ctx context.Context) (fuga.ValuesList, error) {
	res, err := g.db.Fuga.
		Query().All(ctx)

	if err != nil {
		return nil, err
	}

	list := make([]fuga.Values, 0)

	for _, r := range res {
		entity := FugaEntity(*r)
		list = append(list, entity.ToValues())
	}

	return fuga.NewValuesList(list), nil
}

func (g *fugaGateway) Update(ctx context.Context, values fuga.Values) (fuga.Values, error) {
	_, err := g.db.Fuga.
		Update().
		SetName(values.Name().Value()).
		SetNumber(values.Number().Value()).
		SetUpdatedAt(values.UpdatedAt().Value()).
		Save(ctx)

	if err != nil {
		return fuga.Values{}, err
	}

	return values, nil
}

func (g *fugaGateway) Delete(ctx context.Context, id common.Id) error {
	err := g.db.Fuga.DeleteOneID(id.Value()).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
