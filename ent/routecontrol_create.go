// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/routecontrol"
	"github.com/google/uuid"
)

// RouteControlCreate is the builder for creating a RouteControl entity.
type RouteControlCreate struct {
	config
	mutation *RouteControlMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (rcc *RouteControlCreate) SetCreatedAt(t time.Time) *RouteControlCreate {
	rcc.mutation.SetCreatedAt(t)
	return rcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rcc *RouteControlCreate) SetNillableCreatedAt(t *time.Time) *RouteControlCreate {
	if t != nil {
		rcc.SetCreatedAt(*t)
	}
	return rcc
}

// SetUpdatedAt sets the "updated_at" field.
func (rcc *RouteControlCreate) SetUpdatedAt(t time.Time) *RouteControlCreate {
	rcc.mutation.SetUpdatedAt(t)
	return rcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rcc *RouteControlCreate) SetNillableUpdatedAt(t *time.Time) *RouteControlCreate {
	if t != nil {
		rcc.SetUpdatedAt(*t)
	}
	return rcc
}

// SetDistanceMethod sets the "distance_method" field.
func (rcc *RouteControlCreate) SetDistanceMethod(rm routecontrol.DistanceMethod) *RouteControlCreate {
	rcc.mutation.SetDistanceMethod(rm)
	return rcc
}

// SetNillableDistanceMethod sets the "distance_method" field if the given value is not nil.
func (rcc *RouteControlCreate) SetNillableDistanceMethod(rm *routecontrol.DistanceMethod) *RouteControlCreate {
	if rm != nil {
		rcc.SetDistanceMethod(*rm)
	}
	return rcc
}

// SetMileageUnit sets the "mileage_unit" field.
func (rcc *RouteControlCreate) SetMileageUnit(ru routecontrol.MileageUnit) *RouteControlCreate {
	rcc.mutation.SetMileageUnit(ru)
	return rcc
}

// SetNillableMileageUnit sets the "mileage_unit" field if the given value is not nil.
func (rcc *RouteControlCreate) SetNillableMileageUnit(ru *routecontrol.MileageUnit) *RouteControlCreate {
	if ru != nil {
		rcc.SetMileageUnit(*ru)
	}
	return rcc
}

// SetGenerateRoutes sets the "generate_routes" field.
func (rcc *RouteControlCreate) SetGenerateRoutes(b bool) *RouteControlCreate {
	rcc.mutation.SetGenerateRoutes(b)
	return rcc
}

// SetNillableGenerateRoutes sets the "generate_routes" field if the given value is not nil.
func (rcc *RouteControlCreate) SetNillableGenerateRoutes(b *bool) *RouteControlCreate {
	if b != nil {
		rcc.SetGenerateRoutes(*b)
	}
	return rcc
}

// SetID sets the "id" field.
func (rcc *RouteControlCreate) SetID(u uuid.UUID) *RouteControlCreate {
	rcc.mutation.SetID(u)
	return rcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rcc *RouteControlCreate) SetNillableID(u *uuid.UUID) *RouteControlCreate {
	if u != nil {
		rcc.SetID(*u)
	}
	return rcc
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (rcc *RouteControlCreate) SetOrganizationID(id uuid.UUID) *RouteControlCreate {
	rcc.mutation.SetOrganizationID(id)
	return rcc
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (rcc *RouteControlCreate) SetOrganization(o *Organization) *RouteControlCreate {
	return rcc.SetOrganizationID(o.ID)
}

// SetBusinessUnitID sets the "business_unit" edge to the BusinessUnit entity by ID.
func (rcc *RouteControlCreate) SetBusinessUnitID(id uuid.UUID) *RouteControlCreate {
	rcc.mutation.SetBusinessUnitID(id)
	return rcc
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (rcc *RouteControlCreate) SetBusinessUnit(b *BusinessUnit) *RouteControlCreate {
	return rcc.SetBusinessUnitID(b.ID)
}

// Mutation returns the RouteControlMutation object of the builder.
func (rcc *RouteControlCreate) Mutation() *RouteControlMutation {
	return rcc.mutation
}

// Save creates the RouteControl in the database.
func (rcc *RouteControlCreate) Save(ctx context.Context) (*RouteControl, error) {
	rcc.defaults()
	return withHooks(ctx, rcc.sqlSave, rcc.mutation, rcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rcc *RouteControlCreate) SaveX(ctx context.Context) *RouteControl {
	v, err := rcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcc *RouteControlCreate) Exec(ctx context.Context) error {
	_, err := rcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcc *RouteControlCreate) ExecX(ctx context.Context) {
	if err := rcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rcc *RouteControlCreate) defaults() {
	if _, ok := rcc.mutation.CreatedAt(); !ok {
		v := routecontrol.DefaultCreatedAt()
		rcc.mutation.SetCreatedAt(v)
	}
	if _, ok := rcc.mutation.UpdatedAt(); !ok {
		v := routecontrol.DefaultUpdatedAt()
		rcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rcc.mutation.DistanceMethod(); !ok {
		v := routecontrol.DefaultDistanceMethod
		rcc.mutation.SetDistanceMethod(v)
	}
	if _, ok := rcc.mutation.MileageUnit(); !ok {
		v := routecontrol.DefaultMileageUnit
		rcc.mutation.SetMileageUnit(v)
	}
	if _, ok := rcc.mutation.GenerateRoutes(); !ok {
		v := routecontrol.DefaultGenerateRoutes
		rcc.mutation.SetGenerateRoutes(v)
	}
	if _, ok := rcc.mutation.ID(); !ok {
		v := routecontrol.DefaultID()
		rcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rcc *RouteControlCreate) check() error {
	if _, ok := rcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "RouteControl.created_at"`)}
	}
	if _, ok := rcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "RouteControl.updated_at"`)}
	}
	if _, ok := rcc.mutation.DistanceMethod(); !ok {
		return &ValidationError{Name: "distance_method", err: errors.New(`ent: missing required field "RouteControl.distance_method"`)}
	}
	if v, ok := rcc.mutation.DistanceMethod(); ok {
		if err := routecontrol.DistanceMethodValidator(v); err != nil {
			return &ValidationError{Name: "distance_method", err: fmt.Errorf(`ent: validator failed for field "RouteControl.distance_method": %w`, err)}
		}
	}
	if _, ok := rcc.mutation.MileageUnit(); !ok {
		return &ValidationError{Name: "mileage_unit", err: errors.New(`ent: missing required field "RouteControl.mileage_unit"`)}
	}
	if v, ok := rcc.mutation.MileageUnit(); ok {
		if err := routecontrol.MileageUnitValidator(v); err != nil {
			return &ValidationError{Name: "mileage_unit", err: fmt.Errorf(`ent: validator failed for field "RouteControl.mileage_unit": %w`, err)}
		}
	}
	if _, ok := rcc.mutation.GenerateRoutes(); !ok {
		return &ValidationError{Name: "generate_routes", err: errors.New(`ent: missing required field "RouteControl.generate_routes"`)}
	}
	if _, ok := rcc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "RouteControl.organization"`)}
	}
	if _, ok := rcc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit", err: errors.New(`ent: missing required edge "RouteControl.business_unit"`)}
	}
	return nil
}

func (rcc *RouteControlCreate) sqlSave(ctx context.Context) (*RouteControl, error) {
	if err := rcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	rcc.mutation.id = &_node.ID
	rcc.mutation.done = true
	return _node, nil
}

func (rcc *RouteControlCreate) createSpec() (*RouteControl, *sqlgraph.CreateSpec) {
	var (
		_node = &RouteControl{config: rcc.config}
		_spec = sqlgraph.NewCreateSpec(routecontrol.Table, sqlgraph.NewFieldSpec(routecontrol.FieldID, field.TypeUUID))
	)
	if id, ok := rcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rcc.mutation.CreatedAt(); ok {
		_spec.SetField(routecontrol.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rcc.mutation.UpdatedAt(); ok {
		_spec.SetField(routecontrol.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rcc.mutation.DistanceMethod(); ok {
		_spec.SetField(routecontrol.FieldDistanceMethod, field.TypeEnum, value)
		_node.DistanceMethod = value
	}
	if value, ok := rcc.mutation.MileageUnit(); ok {
		_spec.SetField(routecontrol.FieldMileageUnit, field.TypeEnum, value)
		_node.MileageUnit = value
	}
	if value, ok := rcc.mutation.GenerateRoutes(); ok {
		_spec.SetField(routecontrol.FieldGenerateRoutes, field.TypeBool, value)
		_node.GenerateRoutes = value
	}
	if nodes := rcc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   routecontrol.OrganizationTable,
			Columns: []string{routecontrol.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.organization_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rcc.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   routecontrol.BusinessUnitTable,
			Columns: []string{routecontrol.BusinessUnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessunit.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.business_unit_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RouteControlCreateBulk is the builder for creating many RouteControl entities in bulk.
type RouteControlCreateBulk struct {
	config
	err      error
	builders []*RouteControlCreate
}

// Save creates the RouteControl entities in the database.
func (rccb *RouteControlCreateBulk) Save(ctx context.Context) ([]*RouteControl, error) {
	if rccb.err != nil {
		return nil, rccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rccb.builders))
	nodes := make([]*RouteControl, len(rccb.builders))
	mutators := make([]Mutator, len(rccb.builders))
	for i := range rccb.builders {
		func(i int, root context.Context) {
			builder := rccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RouteControlMutation)
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
					_, err = mutators[i+1].Mutate(root, rccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, rccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rccb *RouteControlCreateBulk) SaveX(ctx context.Context) []*RouteControl {
	v, err := rccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rccb *RouteControlCreateBulk) Exec(ctx context.Context) error {
	_, err := rccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rccb *RouteControlCreateBulk) ExecX(ctx context.Context) {
	if err := rccb.Exec(ctx); err != nil {
		panic(err)
	}
}
