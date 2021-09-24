package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Hoge holds the schema definition for the Hoge entity.
type Hoge struct {
	ent.Schema
}

// Fields of the Hoge.
func (Hoge) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("name"),
		field.Int32("number"),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Hoge.
func (Hoge) Edges() []ent.Edge {
	return nil
}
