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

func (g *xxxGateway) Save(ctx context.Context, input *xxx.RepositorySaveInput) (*xxx.RepositorySaveOutput, error) {
	entity := Entity(input.Xxx)

	_, err := g.db.Xxx.
		Create().
		SetID(entity.ID).
		SetName(entity.Name).
		SetNumber(entity.Number).
		SetCreatedAt(entity.CreatedAt).
		SetUpdatedAt(entity.UpdatedAt).
		Save(ctx)

	if err != nil {
		return nil, xe.NewInternalServerError(err, "xxx gateway save failed")
	}

	return &xxx.RepositorySaveOutput{
		Xxx: entity.ToProps(),
	}, nil
}

func (g *xxxGateway) Find(ctx context.Context, input *xxx.RepositoryFindInput) (*xxx.RepositoryFindOutput, error) {
	res, err := g.db.Xxx.Get(ctx, input.Id.Value())

	if err != nil {
		return nil, xe.NewInternalServerError(err, "xxx gateway save failed")
	}

	entity := XxxEntity(*res)

	return &xxx.RepositoryFindOutput{
		Xxx: entity.ToProps(),
	}, nil
}

func (g *xxxGateway) FindAll(ctx context.Context, input *xxx.RepositoryFindAllInput) (*xxx.RepositoryFindAllOutput, error) {
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

	return &xxx.RepositoryFindAllOutput{
		Xxxs: list,
	}, nil
}

func (g *xxxGateway) Update(ctx context.Context, input *xxx.RepositoryUpdateInput) (*xxx.RepositoryUpdateOutput, error) {
	entity := Entity(input.Xxx)

	_, err := g.db.Xxx.
		Update().
		SetName(entity.Name).
		SetNumber(entity.Number).
		SetUpdatedAt(entity.UpdatedAt).
		Save(ctx)

	if err != nil {
		return nil, xe.NewInternalServerError(err, "xxx gateway update failed")
	}

	return &xxx.RepositoryUpdateOutput{
		Xxx: entity.ToProps(),
	}, nil
}

func (g *xxxGateway) Delete(ctx context.Context, input *xxx.RepositoryDeleteInput) (*xxx.RepositoryDeleteOutput, error) {
	err := g.db.Xxx.DeleteOneID(input.Id.Value()).Exec(ctx)

	if err != nil {
		return nil, xe.NewInternalServerError(err, "xxx gateway delete failed")
	}

	return &xxx.RepositoryDeleteOutput{}, nil
}
