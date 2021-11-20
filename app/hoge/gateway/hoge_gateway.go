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

func (g *hogeGateway) Save(ctx context.Context, values hoge.Values) (hoge.Values, error) {
	_, err := g.db.Hoge.
		Create().
		SetID(values.Id().Value()).
		SetName(values.Name().Value()).
		SetNumber(values.Number().Value()).
		SetCreatedAt(values.CreatedAt().Value()).
		SetUpdatedAt(values.UpdatedAt().Value()).
		Save(ctx)

	if err != nil {
		return hoge.Values{}, err
	}

	return values, nil
}

func (g *hogeGateway) Find(ctx context.Context, id common.Id) (hoge.Values, error) {
	res, err := g.db.Hoge.Get(ctx, id.Value())
	if err != nil {
		return hoge.Values{}, err
	}

	entity := HogeEntity(*res)

	return entity.ToValues(), nil
}

func (g *hogeGateway) FindAll(ctx context.Context) (hoge.ValuesList, error) {
	res, err := g.db.Hoge.
		Query().All(ctx)

	if err != nil {
		return nil, err
	}

	list := make([]hoge.Values, 0)

	for _, r := range res {
		entity := HogeEntity(*r)
		list = append(list, entity.ToValues())
	}

	return hoge.NewValuesList(list), nil
}

func (g *hogeGateway) Update(ctx context.Context, values hoge.Values) (hoge.Values, error) {
	_, err := g.db.Hoge.
		Update().
		SetName(values.Name().Value()).
		SetNumber(values.Number().Value()).
		SetUpdatedAt(values.UpdatedAt().Value()).
		Save(ctx)

	if err != nil {
		return hoge.Values{}, err
	}

	return values, nil
}

func (g *hogeGateway) Delete(ctx context.Context, id common.Id) error {
	err := g.db.Hoge.DeleteOneID(id.Value()).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
