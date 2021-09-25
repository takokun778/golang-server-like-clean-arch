package common

import (
	"github.com/google/uuid"
)

type Id uuid.UUID

func (id Id) Value() uuid.UUID {
	return uuid.UUID(id)
}

func ParseId(value string) (Id, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return Id{}, err
	}
	return Id(id), nil
}

func CreateRandomId() Id {
	id, _ := uuid.NewRandom()
	return Id(id)
}
