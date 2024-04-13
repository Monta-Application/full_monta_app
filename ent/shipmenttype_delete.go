// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/emoss08/trenova/ent/shipmenttype"
)

// ShipmentTypeDelete is the builder for deleting a ShipmentType entity.
type ShipmentTypeDelete struct {
	config
	hooks    []Hook
	mutation *ShipmentTypeMutation
}

// Where appends a list predicates to the ShipmentTypeDelete builder.
func (std *ShipmentTypeDelete) Where(ps ...predicate.ShipmentType) *ShipmentTypeDelete {
	std.mutation.Where(ps...)
	return std
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (std *ShipmentTypeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, std.sqlExec, std.mutation, std.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (std *ShipmentTypeDelete) ExecX(ctx context.Context) int {
	n, err := std.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (std *ShipmentTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(shipmenttype.Table, sqlgraph.NewFieldSpec(shipmenttype.FieldID, field.TypeUUID))
	if ps := std.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, std.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	std.mutation.done = true
	return affected, err
}

// ShipmentTypeDeleteOne is the builder for deleting a single ShipmentType entity.
type ShipmentTypeDeleteOne struct {
	std *ShipmentTypeDelete
}

// Where appends a list predicates to the ShipmentTypeDelete builder.
func (stdo *ShipmentTypeDeleteOne) Where(ps ...predicate.ShipmentType) *ShipmentTypeDeleteOne {
	stdo.std.mutation.Where(ps...)
	return stdo
}

// Exec executes the deletion query.
func (stdo *ShipmentTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := stdo.std.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{shipmenttype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (stdo *ShipmentTypeDeleteOne) ExecX(ctx context.Context) {
	if err := stdo.Exec(ctx); err != nil {
		panic(err)
	}
}
