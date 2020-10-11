package schema

import (
    "github.com/facebook/ent"
    "github.com/facebook/ent/schema/field"
    "github.com/facebook/ent/schema/edge"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("DoctorID"),
		field.String("DoctorEmail"),
		field.String("DoctorName"),
		//field.Time("Date"),
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
