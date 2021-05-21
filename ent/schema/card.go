package schema

import (
	"encoding/json"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("content", json.RawMessage{}),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("card").
			Unique(),
	}
}
