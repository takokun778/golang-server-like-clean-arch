package gateway

import (
	"context"
	"xxx/app/common/gateway"
	"xxx/app/domain/common"
	"xxx/app/domain/xxx"
)

type xxxGateway struct {
	db *gateway.Database
}

func NewXxxGateway() xxx.Repository {
	gw := new(xxxGateway)
	gw.db = gateway.DatabaseConnect()
	return gw
}

func (g *xxxGateway) Save(ctx context.Context, values xxx.Values) (xxx.Values, error) {
	_, err := g.db.Xxx.
		Create().
		SetID(values.Id().Value()).
		SetName(values.Name().Value()).
		SetNumber(values.Number().Value()).
		SetCreatedAt(values.CreatedAt().Value()).
		SetUpdatedAt(values.UpdatedAt().Value()).
		Save(ctx)

	if err != nil {
		return xxx.Values{}, err
	}

	return values, nil
}

func (g *xxxGateway) Find(ctx context.Context, id common.Id) (xxx.Values, error) {
	res, err := g.db.Xxx.Get(ctx, id.Value())

	if err != nil {
		return xxx.Values{}, err
	}

	entity := XxxEntity(*res)

	return entity.ToValues(), nil
}

func (g *xxxGateway) FindAll(ctx context.Context) ([]xxx.Values, error) {
	res, err := g.db.Xxx.
		Query().All(ctx)

	if err != nil {
		return nil, err
	}

	list := make([]xxx.Values, 0)

	for _, r := range res {
		entity := XxxEntity(*r)
		list = append(list, entity.ToValues())
	}

	return list, nil
}

func (g *xxxGateway) Update(ctx context.Context, values xxx.Values) (xxx.Values, error) {
	_, err := g.db.Xxx.
		Update().
		SetName(values.Name().Value()).
		SetNumber(values.Number().Value()).
		SetUpdatedAt(values.UpdatedAt().Value()).
		Save(ctx)

	if err != nil {
		return xxx.Values{}, err
	}

	return values, nil
}

func (g *xxxGateway) Delete(ctx context.Context, id common.Id) error {
	err := g.db.Xxx.DeleteOneID(id.Value()).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
