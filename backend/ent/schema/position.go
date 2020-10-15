package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent/schema/edge"
)

// Position holds the schema definition for the Position entity.
type Position struct {
	ent.Schema
}

// Fields of the Position.
func (Position) Fields() []ent.Field {
	return []ent.Field{
		field.String("PositionName").Unique(),

	}
}

// Edges of the Position.
func (Position) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("PositionUser",User.Type).StorageKey(edge.Column("PositionID")),
	}
}
