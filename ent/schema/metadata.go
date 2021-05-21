package schema

import "entgo.io/ent"

// Metadata holds the schema definition for the Metadata entity.
type Metadata struct {
	ent.Schema
}

// Fields of the Metadata.
func (Metadata) Fields() []ent.Field {
	return nil
}

// Edges of the Metadata.
func (Metadata) Edges() []ent.Edge {
	return nil
}
