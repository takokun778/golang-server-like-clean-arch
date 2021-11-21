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

func (g *xxxGateway) Save(ctx context.Context, props xxx.Props) (xxx.Props, error) {
	_, err := g.db.Xxx.
		Create().
		SetID(props.Id().Value()).
		SetName(props.Name().Value()).
		SetNumber(props.Number().Value()).
		SetCreatedAt(props.CreatedAt().Value()).
		SetUpdatedAt(props.UpdatedAt().Value()).
		Save(ctx)

	if err != nil {
		return xxx.Props{}, common.NewInternalServerError(err, "xxx gateway save failed")
	}

	return props, nil
}

func (g *xxxGateway) Find(ctx context.Context, id common.Id) (xxx.Props, error) {
	res, err := g.db.Xxx.Get(ctx, id.Value())

	if err != nil {
		return xxx.Props{}, common.NewInternalServerError(err, "xxx gateway save failed")
	}

	entity := XxxEntity(*res)

	return entity.ToProps(), nil
}

func (g *xxxGateway) FindAll(ctx context.Context) ([]xxx.Props, error) {
	res, err := g.db.Xxx.
		Query().All(ctx)

	if err != nil {
		return nil, common.NewInternalServerError(err, "xxx gateway find all failed")
	}

	list := make([]xxx.Props, 0)

	for _, r := range res {
		entity := XxxEntity(*r)
		list = append(list, entity.ToProps())
	}

	return list, nil
}

func (g *xxxGateway) Update(ctx context.Context, props xxx.Props) (xxx.Props, error) {
	_, err := g.db.Xxx.
		Update().
		SetName(props.Name().Value()).
		SetNumber(props.Number().Value()).
		SetUpdatedAt(props.UpdatedAt().Value()).
		Save(ctx)

	if err != nil {
		return xxx.Props{}, common.NewInternalServerError(err, "xxx gateway update failed")
	}

	return props, nil
}

func (g *xxxGateway) Delete(ctx context.Context, id common.Id) error {
	err := g.db.Xxx.DeleteOneID(id.Value()).Exec(ctx)

	if err != nil {
		return common.NewInternalServerError(err, "xxx gateway delete failed")
	}

	return nil
}
