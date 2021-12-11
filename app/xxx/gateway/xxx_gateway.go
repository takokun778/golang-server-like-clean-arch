package gateway

import (
	"context"

	dErr "xxx/app/domain/error"
	"xxx/app/domain/xxx"
)

type xxxGateway struct {
	sql XxxSql
}

func NewXxxGateway(sql XxxSql) xxx.Repository {
	gw := new(xxxGateway)
	gw.sql = sql
	return gw
}

func (g *xxxGateway) Save(ctx context.Context, item *xxx.RepositorySaveItem) (xxx.Props, error) {
	entity := Entity(item.Xxx)

	err := g.sql.InsertXxx(ctx, entity)

	if err != nil {
		return xxx.Props{}, dErr.NewInternalServerError(err, "xxx gateway save failed")
	}

	return entity.ToProps(), nil
}

func (g *xxxGateway) Find(ctx context.Context, item *xxx.RepositoryFindItem) (xxx.Props, error) {
	res, err := g.sql.SelectXxx(ctx, item.Id.Value())

	if err != nil {
		return xxx.Props{}, dErr.NewInternalServerError(err, "xxx gateway save failed")
	}

	entity := XxxEntity(res)

	return entity.ToProps(), nil
}

func (g *xxxGateway) FindAll(ctx context.Context, item *xxx.RepositoryFindAllItem) ([]xxx.Props, error) {
	res, err := g.sql.SelectXxxAll(ctx)

	if err != nil {
		return nil, dErr.NewInternalServerError(err, "xxx gateway find all failed")
	}

	list := make([]xxx.Props, 0)

	for _, r := range res {
		entity := XxxEntity(r)
		list = append(list, entity.ToProps())
	}

	return list, nil
}

func (g *xxxGateway) Update(ctx context.Context, item *xxx.RepositoryUpdateItem) (xxx.Props, error) {
	entity := Entity(item.Xxx)

	err := g.sql.UpdateXxx(ctx, entity)

	if err != nil {
		return xxx.Props{}, dErr.NewInternalServerError(err, "xxx gateway update failed")
	}

	return entity.ToProps(), nil
}

func (g *xxxGateway) Delete(ctx context.Context, item *xxx.RepositoryDeleteItem) error {
	err := g.sql.DeleteXxx(ctx, item.Id.Value())

	if err != nil {
		return dErr.NewInternalServerError(err, "xxx gateway delete failed")
	}

	return nil
}
