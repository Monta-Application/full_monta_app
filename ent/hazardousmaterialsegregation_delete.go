// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/hazardousmaterialsegregation"
	"github.com/emoss08/trenova/ent/predicate"
)

// HazardousMaterialSegregationDelete is the builder for deleting a HazardousMaterialSegregation entity.
type HazardousMaterialSegregationDelete struct {
	config
	hooks    []Hook
	mutation *HazardousMaterialSegregationMutation
}

// Where appends a list predicates to the HazardousMaterialSegregationDelete builder.
func (hmsd *HazardousMaterialSegregationDelete) Where(ps ...predicate.HazardousMaterialSegregation) *HazardousMaterialSegregationDelete {
	hmsd.mutation.Where(ps...)
	return hmsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hmsd *HazardousMaterialSegregationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, hmsd.sqlExec, hmsd.mutation, hmsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (hmsd *HazardousMaterialSegregationDelete) ExecX(ctx context.Context) int {
	n, err := hmsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hmsd *HazardousMaterialSegregationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(hazardousmaterialsegregation.Table, sqlgraph.NewFieldSpec(hazardousmaterialsegregation.FieldID, field.TypeUUID))
	if ps := hmsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, hmsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	hmsd.mutation.done = true
	return affected, err
}

// HazardousMaterialSegregationDeleteOne is the builder for deleting a single HazardousMaterialSegregation entity.
type HazardousMaterialSegregationDeleteOne struct {
	hmsd *HazardousMaterialSegregationDelete
}

// Where appends a list predicates to the HazardousMaterialSegregationDelete builder.
func (hmsdo *HazardousMaterialSegregationDeleteOne) Where(ps ...predicate.HazardousMaterialSegregation) *HazardousMaterialSegregationDeleteOne {
	hmsdo.hmsd.mutation.Where(ps...)
	return hmsdo
}

// Exec executes the deletion query.
func (hmsdo *HazardousMaterialSegregationDeleteOne) Exec(ctx context.Context) error {
	n, err := hmsdo.hmsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{hazardousmaterialsegregation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hmsdo *HazardousMaterialSegregationDeleteOne) ExecX(ctx context.Context) {
	if err := hmsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
