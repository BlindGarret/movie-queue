package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Rank holds the schema definition for the Rank entity.
type Rank struct {
	ent.Schema
}

// Fields of the Rank.
func (Rank) Fields() []ent.Field {
	return []ent.Field{
		field.String("next").Default("0|AA"),
	}
}

// Edges of the Rank.
func (Rank) Edges() []ent.Edge {
	return nil
}
