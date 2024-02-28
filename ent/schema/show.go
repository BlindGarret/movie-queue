package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Show holds the schema definition for the Show entity.
type Show struct {
	ent.Schema
}

// Fields of the Show.
func (Show) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("order"),
		field.Enum("type").Values("movie", "tv"),
		field.Time("addedDate").Default(time.Now),
	}
}

// Edges of the Show.
func (Show) Edges() []ent.Edge {
	return nil
}
