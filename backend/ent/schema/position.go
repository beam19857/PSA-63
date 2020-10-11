package schema

import (
    "github.com/facebook/ent"
    "github.com/facebook/ent/schema/field"
    "github.com/facebook/ent/schema/edge"
)

// Position holds the schema definition for the Position entity.
type Position struct {
	ent.Schema
}

// Fields of the Position.
func (Position) Fields() []ent.Field {
	return []ent.Field{
		field.Int("PositionID"),
		field.String("PositionName"),

	}
}

// Edges of the Position.
func (Position) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("PositionUser",User.Type),
	}
}
