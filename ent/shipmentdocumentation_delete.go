// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/emoss08/trenova/ent/shipmentdocumentation"
)

// ShipmentDocumentationDelete is the builder for deleting a ShipmentDocumentation entity.
type ShipmentDocumentationDelete struct {
	config
	hooks    []Hook
	mutation *ShipmentDocumentationMutation
}

// Where appends a list predicates to the ShipmentDocumentationDelete builder.
func (sdd *ShipmentDocumentationDelete) Where(ps ...predicate.ShipmentDocumentation) *ShipmentDocumentationDelete {
	sdd.mutation.Where(ps...)
	return sdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sdd *ShipmentDocumentationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sdd.sqlExec, sdd.mutation, sdd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sdd *ShipmentDocumentationDelete) ExecX(ctx context.Context) int {
	n, err := sdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sdd *ShipmentDocumentationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(shipmentdocumentation.Table, sqlgraph.NewFieldSpec(shipmentdocumentation.FieldID, field.TypeUUID))
	if ps := sdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sdd.mutation.done = true
	return affected, err
}

// ShipmentDocumentationDeleteOne is the builder for deleting a single ShipmentDocumentation entity.
type ShipmentDocumentationDeleteOne struct {
	sdd *ShipmentDocumentationDelete
}

// Where appends a list predicates to the ShipmentDocumentationDelete builder.
func (sddo *ShipmentDocumentationDeleteOne) Where(ps ...predicate.ShipmentDocumentation) *ShipmentDocumentationDeleteOne {
	sddo.sdd.mutation.Where(ps...)
	return sddo
}

// Exec executes the deletion query.
func (sddo *ShipmentDocumentationDeleteOne) Exec(ctx context.Context) error {
	n, err := sddo.sdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{shipmentdocumentation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sddo *ShipmentDocumentationDeleteOne) ExecX(ctx context.Context) {
	if err := sddo.Exec(ctx); err != nil {
		panic(err)
	}
}
