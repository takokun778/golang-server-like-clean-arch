package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Xxx holds the schema definition for the Xxx entity.
type Xxx struct {
	ent.Schema
}

// Fields of the Xxx.
func (Xxx) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("name").Optional(),
		field.Int("number").Optional(),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Xxx.
func (Xxx) Edges() []ent.Edge {
	return nil
}
