package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Fuga holds the schema definition for the Fuga entity.
type Fuga struct {
	ent.Schema
}

// Fields of the Fuga.
func (Fuga) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("name"),
		field.Int32("number"),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Fuga.
func (Fuga) Edges() []ent.Edge {
	return nil
}
