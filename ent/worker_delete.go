// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/emoss08/trenova/ent/worker"
)

// WorkerDelete is the builder for deleting a Worker entity.
type WorkerDelete struct {
	config
	hooks    []Hook
	mutation *WorkerMutation
}

// Where appends a list predicates to the WorkerDelete builder.
func (wd *WorkerDelete) Where(ps ...predicate.Worker) *WorkerDelete {
	wd.mutation.Where(ps...)
	return wd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wd *WorkerDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wd.sqlExec, wd.mutation, wd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wd *WorkerDelete) ExecX(ctx context.Context) int {
	n, err := wd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wd *WorkerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(worker.Table, sqlgraph.NewFieldSpec(worker.FieldID, field.TypeUUID))
	if ps := wd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wd.mutation.done = true
	return affected, err
}

// WorkerDeleteOne is the builder for deleting a single Worker entity.
type WorkerDeleteOne struct {
	wd *WorkerDelete
}

// Where appends a list predicates to the WorkerDelete builder.
func (wdo *WorkerDeleteOne) Where(ps ...predicate.Worker) *WorkerDeleteOne {
	wdo.wd.mutation.Where(ps...)
	return wdo
}

// Exec executes the deletion query.
func (wdo *WorkerDeleteOne) Exec(ctx context.Context) error {
	n, err := wdo.wd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{worker.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wdo *WorkerDeleteOne) ExecX(ctx context.Context) {
	if err := wdo.Exec(ctx); err != nil {
		panic(err)
	}
}
