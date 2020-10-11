package schema

import (
    "github.com/facebook/ent"
    "github.com/facebook/ent/schema/field"
    "github.com/facebook/ent/schema/edge"
)


// Expertise holds the schema definition for the Expertise entity.
type Expertise struct {
	ent.Schema
}

// Fields of the Expertise.
func (Expertise) Fields() []ent.Field {
	return []ent.Field{
		field.Int("ExpertiseID"),
		field.String("ExpertiseName"),
		//field.String("Licenes"),
	}
}

// Edges of the Expertise.
func (Expertise) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ExpertiseUser",User.Type),
	}
}
