// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/beam19857/app/ent/department"
	"github.com/beam19857/app/ent/expertise"
	"github.com/beam19857/app/ent/position"
	"github.com/beam19857/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetDoctorName sets the DoctorName field.
func (uc *UserCreate) SetDoctorName(s string) *UserCreate {
	uc.mutation.SetDoctorName(s)
	return uc
}

// SetDoctorEmail sets the DoctorEmail field.
func (uc *UserCreate) SetDoctorEmail(s string) *UserCreate {
	uc.mutation.SetDoctorEmail(s)
	return uc
}

// SetUserDepartmentID sets the UserDepartment edge to Department by id.
func (uc *UserCreate) SetUserDepartmentID(id int) *UserCreate {
	uc.mutation.SetUserDepartmentID(id)
	return uc
}

// SetNillableUserDepartmentID sets the UserDepartment edge to Department by id if the given value is not nil.
func (uc *UserCreate) SetNillableUserDepartmentID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetUserDepartmentID(*id)
	}
	return uc
}

// SetUserDepartment sets the UserDepartment edge to Department.
func (uc *UserCreate) SetUserDepartment(d *Department) *UserCreate {
	return uc.SetUserDepartmentID(d.ID)
}

// SetUserExpertiseID sets the UserExpertise edge to Expertise by id.
func (uc *UserCreate) SetUserExpertiseID(id int) *UserCreate {
	uc.mutation.SetUserExpertiseID(id)
	return uc
}

// SetNillableUserExpertiseID sets the UserExpertise edge to Expertise by id if the given value is not nil.
func (uc *UserCreate) SetNillableUserExpertiseID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetUserExpertiseID(*id)
	}
	return uc
}

// SetUserExpertise sets the UserExpertise edge to Expertise.
func (uc *UserCreate) SetUserExpertise(e *Expertise) *UserCreate {
	return uc.SetUserExpertiseID(e.ID)
}

// SetUserPositionID sets the UserPosition edge to Position by id.
func (uc *UserCreate) SetUserPositionID(id int) *UserCreate {
	uc.mutation.SetUserPositionID(id)
	return uc
}

// SetNillableUserPositionID sets the UserPosition edge to Position by id if the given value is not nil.
func (uc *UserCreate) SetNillableUserPositionID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetUserPositionID(*id)
	}
	return uc
}

// SetUserPosition sets the UserPosition edge to Position.
func (uc *UserCreate) SetUserPosition(p *Position) *UserCreate {
	return uc.SetUserPositionID(p.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if _, ok := uc.mutation.DoctorName(); !ok {
		return nil, &ValidationError{Name: "DoctorName", err: errors.New("ent: missing required field \"DoctorName\"")}
	}
	if _, ok := uc.mutation.DoctorEmail(); !ok {
		return nil, &ValidationError{Name: "DoctorEmail", err: errors.New("ent: missing required field \"DoctorEmail\"")}
	}
	var (
		err  error
		node *User
	)
	if len(uc.hooks) == 0 {
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	u, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	u.ID = int(id)
	return u, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		u     = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.DoctorName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDoctorName,
		})
		u.DoctorName = value
	}
	if value, ok := uc.mutation.DoctorEmail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDoctorEmail,
		})
		u.DoctorEmail = value
	}
	if nodes := uc.mutation.UserDepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.UserDepartmentTable,
			Columns: []string{user.UserDepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserExpertiseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.UserExpertiseTable,
			Columns: []string{user.UserExpertiseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: expertise.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserPositionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.UserPositionTable,
			Columns: []string{user.UserPositionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: position.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return u, _spec
}
