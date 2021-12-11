package ent

import (
	"context"

	"xxx/app/xxx/gateway"

	"github.com/google/uuid"
)

func (c *Connector) InsertXxx(ctx context.Context, xxx gateway.XxxEntity) error {
	_, err := c.Xxx.
		Create().
		SetID(xxx.Id).
		SetName(xxx.Name).
		SetNumber(xxx.Number).
		SetCreatedAt(xxx.CreatedAt).
		SetUpdatedAt(xxx.UpdatedAt).
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (c *Connector) SelectXxx(ctx context.Context, id uuid.UUID) (gateway.XxxEntity, error) {
	res, err := c.Xxx.Get(ctx, id)

	if err != nil {
		return gateway.XxxEntity{}, err
	}

	entity := gateway.XxxEntity{
		Id:        res.ID,
		Name:      res.Name,
		Number:    res.Number,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}

	return entity, nil
}

func (c *Connector) SelectXxxAll(ctx context.Context) ([]gateway.XxxEntity, error) {
	res, err := c.Xxx.
		Query().All(ctx)

	if err != nil {
		return nil, err
	}

	entities := make([]gateway.XxxEntity, 0)

	for _, r := range res {
		entity := gateway.XxxEntity{
			Id:        r.ID,
			Name:      r.Name,
			Number:    r.Number,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

func (c *Connector) UpdateXxx(ctx context.Context, xxx gateway.XxxEntity) error {
	_, err := c.Xxx.
		Update().
		SetName(xxx.Name).
		SetNumber(xxx.Number).
		SetUpdatedAt(xxx.UpdatedAt).
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (c *Connector) DeleteXxx(ctx context.Context, id uuid.UUID) error {
	err := c.Xxx.DeleteOneID(id).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
