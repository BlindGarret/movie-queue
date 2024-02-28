// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BlindGarret/movie-queue/ent/show"
)

// ShowCreate is the builder for creating a Show entity.
type ShowCreate struct {
	config
	mutation *ShowMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sc *ShowCreate) SetName(s string) *ShowCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetOrder sets the "order" field.
func (sc *ShowCreate) SetOrder(s string) *ShowCreate {
	sc.mutation.SetOrder(s)
	return sc
}

// SetType sets the "type" field.
func (sc *ShowCreate) SetType(s show.Type) *ShowCreate {
	sc.mutation.SetType(s)
	return sc
}

// SetAddedDate sets the "addedDate" field.
func (sc *ShowCreate) SetAddedDate(t time.Time) *ShowCreate {
	sc.mutation.SetAddedDate(t)
	return sc
}

// SetNillableAddedDate sets the "addedDate" field if the given value is not nil.
func (sc *ShowCreate) SetNillableAddedDate(t *time.Time) *ShowCreate {
	if t != nil {
		sc.SetAddedDate(*t)
	}
	return sc
}

// Mutation returns the ShowMutation object of the builder.
func (sc *ShowCreate) Mutation() *ShowMutation {
	return sc.mutation
}

// Save creates the Show in the database.
func (sc *ShowCreate) Save(ctx context.Context) (*Show, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ShowCreate) SaveX(ctx context.Context) *Show {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ShowCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ShowCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ShowCreate) defaults() {
	if _, ok := sc.mutation.AddedDate(); !ok {
		v := show.DefaultAddedDate()
		sc.mutation.SetAddedDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ShowCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Show.name"`)}
	}
	if _, ok := sc.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "Show.order"`)}
	}
	if _, ok := sc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Show.type"`)}
	}
	if v, ok := sc.mutation.GetType(); ok {
		if err := show.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Show.type": %w`, err)}
		}
	}
	if _, ok := sc.mutation.AddedDate(); !ok {
		return &ValidationError{Name: "addedDate", err: errors.New(`ent: missing required field "Show.addedDate"`)}
	}
	return nil
}

func (sc *ShowCreate) sqlSave(ctx context.Context) (*Show, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *ShowCreate) createSpec() (*Show, *sqlgraph.CreateSpec) {
	var (
		_node = &Show{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(show.Table, sqlgraph.NewFieldSpec(show.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(show.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Order(); ok {
		_spec.SetField(show.FieldOrder, field.TypeString, value)
		_node.Order = value
	}
	if value, ok := sc.mutation.GetType(); ok {
		_spec.SetField(show.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := sc.mutation.AddedDate(); ok {
		_spec.SetField(show.FieldAddedDate, field.TypeTime, value)
		_node.AddedDate = value
	}
	return _node, _spec
}

// ShowCreateBulk is the builder for creating many Show entities in bulk.
type ShowCreateBulk struct {
	config
	err      error
	builders []*ShowCreate
}

// Save creates the Show entities in the database.
func (scb *ShowCreateBulk) Save(ctx context.Context) ([]*Show, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Show, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShowMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ShowCreateBulk) SaveX(ctx context.Context) []*Show {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ShowCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ShowCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}