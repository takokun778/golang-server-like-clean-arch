package gateway

import (
	"context"
	"xxx/app/common/gateway"
	xe "xxx/app/domain/error"
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

func (g *xxxGateway) Save(ctx context.Context, item *xxx.RepositorySaveItem) (xxx.Props, error) {
	entity := Entity(item.Xxx)

	_, err := g.db.Xxx.
		Create().
		SetID(entity.ID).
		SetName(entity.Name).
		SetNumber(entity.Number).
		SetCreatedAt(entity.CreatedAt).
		SetUpdatedAt(entity.UpdatedAt).
		Save(ctx)

	if err != nil {
		return xxx.Props{}, xe.NewInternalServerError(err, "xxx gateway save failed")
	}

	return entity.ToProps(), nil
}

func (g *xxxGateway) Find(ctx context.Context, item *xxx.RepositoryFindItem) (xxx.Props, error) {
	res, err := g.db.Xxx.Get(ctx, item.Id.Value())

	if err != nil {
		return xxx.Props{}, xe.NewInternalServerError(err, "xxx gateway save failed")
	}

	entity := XxxEntity(*res)

	return entity.ToProps(), nil
}

func (g *xxxGateway) FindAll(ctx context.Context, item *xxx.RepositoryFindAllItem) ([]xxx.Props, error) {
	res, err := g.db.Xxx.
		Query().All(ctx)

	if err != nil {
		return nil, xe.NewInternalServerError(err, "xxx gateway find all failed")
	}

	list := make([]xxx.Props, 0)

	for _, r := range res {
		entity := XxxEntity(*r)
		list = append(list, entity.ToProps())
	}

	return list, nil
}

func (g *xxxGateway) Update(ctx context.Context, item *xxx.RepositoryUpdateItem) (xxx.Props, error) {
	entity := Entity(item.Xxx)

	_, err := g.db.Xxx.
		Update().
		SetName(entity.Name).
		SetNumber(entity.Number).
		SetUpdatedAt(entity.UpdatedAt).
		Save(ctx)

	if err != nil {
		return xxx.Props{}, xe.NewInternalServerError(err, "xxx gateway update failed")
	}

	return entity.ToProps(), nil
}

func (g *xxxGateway) Delete(ctx context.Context, item *xxx.RepositoryDeleteItem) error {
	err := g.db.Xxx.DeleteOneID(item.Id.Value()).Exec(ctx)

	if err != nil {
		return xe.NewInternalServerError(err, "xxx gateway delete failed")
	}

	return nil
}
