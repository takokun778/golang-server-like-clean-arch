package xxx

import (
	"github.com/google/uuid"
)

type id uuid.UUID

func (i id) Value() uuid.UUID {
	return uuid.UUID(i)
}

func ParseId(value string) (id, error) {
	i, err := uuid.Parse(value)

	if err != nil {
		return id{}, err
	}

	return id(i), nil
}

func CreateRandomId() id {
	i, _ := uuid.NewRandom()
	return id(i)
}
