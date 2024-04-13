// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/generalledgeraccount"
	"github.com/emoss08/trenova/ent/predicate"
)

// GeneralLedgerAccountDelete is the builder for deleting a GeneralLedgerAccount entity.
type GeneralLedgerAccountDelete struct {
	config
	hooks    []Hook
	mutation *GeneralLedgerAccountMutation
}

// Where appends a list predicates to the GeneralLedgerAccountDelete builder.
func (glad *GeneralLedgerAccountDelete) Where(ps ...predicate.GeneralLedgerAccount) *GeneralLedgerAccountDelete {
	glad.mutation.Where(ps...)
	return glad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (glad *GeneralLedgerAccountDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, glad.sqlExec, glad.mutation, glad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (glad *GeneralLedgerAccountDelete) ExecX(ctx context.Context) int {
	n, err := glad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (glad *GeneralLedgerAccountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(generalledgeraccount.Table, sqlgraph.NewFieldSpec(generalledgeraccount.FieldID, field.TypeUUID))
	if ps := glad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, glad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	glad.mutation.done = true
	return affected, err
}

// GeneralLedgerAccountDeleteOne is the builder for deleting a single GeneralLedgerAccount entity.
type GeneralLedgerAccountDeleteOne struct {
	glad *GeneralLedgerAccountDelete
}

// Where appends a list predicates to the GeneralLedgerAccountDelete builder.
func (glado *GeneralLedgerAccountDeleteOne) Where(ps ...predicate.GeneralLedgerAccount) *GeneralLedgerAccountDeleteOne {
	glado.glad.mutation.Where(ps...)
	return glado
}

// Exec executes the deletion query.
func (glado *GeneralLedgerAccountDeleteOne) Exec(ctx context.Context) error {
	n, err := glado.glad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{generalledgeraccount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (glado *GeneralLedgerAccountDeleteOne) ExecX(ctx context.Context) {
	if err := glado.Exec(ctx); err != nil {
		panic(err)
	}
}
