package schema

import "entgo.io/ent"

// Info holds the schema definition for the Info entity.
type Info struct {
	ent.Schema
}

// Fields of the Info.
func (Info) Fields() []ent.Field {
	return nil
}

// Edges of the Info.
func (Info) Edges() []ent.Edge {
	return nil
}
