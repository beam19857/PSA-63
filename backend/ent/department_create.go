// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/beam19857/app/ent/department"
	"github.com/beam19857/app/ent/user"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// DepartmentCreate is the builder for creating a Department entity.
type DepartmentCreate struct {
	config
	mutation *DepartmentMutation
	hooks    []Hook
}

// SetDepartmentID sets the DepartmentID field.
func (dc *DepartmentCreate) SetDepartmentID(i int) *DepartmentCreate {
	dc.mutation.SetDepartmentID(i)
	return dc
}

// SetDepartmentName sets the DepartmentName field.
func (dc *DepartmentCreate) SetDepartmentName(s string) *DepartmentCreate {
	dc.mutation.SetDepartmentName(s)
	return dc
}

// AddDepartmentUserIDs adds the DepartmentUser edge to User by ids.
func (dc *DepartmentCreate) AddDepartmentUserIDs(ids ...int) *DepartmentCreate {
	dc.mutation.AddDepartmentUserIDs(ids...)
	return dc
}

// AddDepartmentUser adds the DepartmentUser edges to User.
func (dc *DepartmentCreate) AddDepartmentUser(u ...*User) *DepartmentCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return dc.AddDepartmentUserIDs(ids...)
}

// Mutation returns the DepartmentMutation object of the builder.
func (dc *DepartmentCreate) Mutation() *DepartmentMutation {
	return dc.mutation
}

// Save creates the Department in the database.
func (dc *DepartmentCreate) Save(ctx context.Context) (*Department, error) {
	if err := dc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Department
	)
	if len(dc.hooks) == 0 {
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DepartmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DepartmentCreate) SaveX(ctx context.Context) *Department {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dc *DepartmentCreate) preSave() error {
	if _, ok := dc.mutation.DepartmentID(); !ok {
		return &ValidationError{Name: "DepartmentID", err: errors.New("ent: missing required field \"DepartmentID\"")}
	}
	if _, ok := dc.mutation.DepartmentName(); !ok {
		return &ValidationError{Name: "DepartmentName", err: errors.New("ent: missing required field \"DepartmentName\"")}
	}
	return nil
}

func (dc *DepartmentCreate) sqlSave(ctx context.Context) (*Department, error) {
	d, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	d.ID = int(id)
	return d, nil
}

func (dc *DepartmentCreate) createSpec() (*Department, *sqlgraph.CreateSpec) {
	var (
		d     = &Department{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: department.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: department.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.DepartmentID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: department.FieldDepartmentID,
		})
		d.DepartmentID = value
	}
	if value, ok := dc.mutation.DepartmentName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: department.FieldDepartmentName,
		})
		d.DepartmentName = value
	}
	if nodes := dc.mutation.DepartmentUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.DepartmentUserTable,
			Columns: []string{department.DepartmentUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return d, _spec
}

// DepartmentCreateBulk is the builder for creating a bulk of Department entities.
type DepartmentCreateBulk struct {
	config
	builders []*DepartmentCreate
}

// Save creates the Department entities in the database.
func (dcb *DepartmentCreateBulk) Save(ctx context.Context) ([]*Department, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Department, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*DepartmentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (dcb *DepartmentCreateBulk) SaveX(ctx context.Context) []*Department {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}