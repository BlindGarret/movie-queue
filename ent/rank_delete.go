// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BlindGarret/movie-queue/ent/predicate"
	"github.com/BlindGarret/movie-queue/ent/rank"
)

// RankDelete is the builder for deleting a Rank entity.
type RankDelete struct {
	config
	hooks    []Hook
	mutation *RankMutation
}

// Where appends a list predicates to the RankDelete builder.
func (rd *RankDelete) Where(ps ...predicate.Rank) *RankDelete {
	rd.mutation.Where(ps...)
	return rd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rd *RankDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rd.sqlExec, rd.mutation, rd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rd *RankDelete) ExecX(ctx context.Context) int {
	n, err := rd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rd *RankDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(rank.Table, sqlgraph.NewFieldSpec(rank.FieldID, field.TypeInt))
	if ps := rd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rd.mutation.done = true
	return affected, err
}

// RankDeleteOne is the builder for deleting a single Rank entity.
type RankDeleteOne struct {
	rd *RankDelete
}

// Where appends a list predicates to the RankDelete builder.
func (rdo *RankDeleteOne) Where(ps ...predicate.Rank) *RankDeleteOne {
	rdo.rd.mutation.Where(ps...)
	return rdo
}

// Exec executes the deletion query.
func (rdo *RankDeleteOne) Exec(ctx context.Context) error {
	n, err := rdo.rd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{rank.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdo *RankDeleteOne) ExecX(ctx context.Context) {
	if err := rdo.Exec(ctx); err != nil {
		panic(err)
	}
}
