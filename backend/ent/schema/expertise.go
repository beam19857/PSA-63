package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent/schema/edge"
)


// Expertise holds the schema definition for the Expertise entity.
type Expertise struct {
	ent.Schema
}

// Fields of the Expertise.
func (Expertise) Fields() []ent.Field {
	return []ent.Field{
		field.String("ExpertiseName"),
	}
}

// Edges of the Expertise.
func (Expertise) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ExpertiseUser",User.Type).StorageKey(edge.Column("ExpertiseID")),
	}
}
