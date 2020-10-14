package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent/schema/edge"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("DoctorName"),
		field.String("DoctorEmail"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("UserDepartment",Department.Type).
			Ref("DepartmentUser").Unique(),
		edge.From("UserExpertise",Expertise.Type).
			Ref("ExpertiseUser").Unique(),
		edge.From("UserPosition",Position.Type).
			Ref("PositionUser").Unique(),
	}
}
