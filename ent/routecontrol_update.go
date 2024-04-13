// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/emoss08/trenova/ent/routecontrol"
	"github.com/google/uuid"
)

// RouteControlUpdate is the builder for updating RouteControl entities.
type RouteControlUpdate struct {
	config
	hooks     []Hook
	mutation  *RouteControlMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the RouteControlUpdate builder.
func (rcu *RouteControlUpdate) Where(ps ...predicate.RouteControl) *RouteControlUpdate {
	rcu.mutation.Where(ps...)
	return rcu
}

// SetUpdatedAt sets the "updated_at" field.
func (rcu *RouteControlUpdate) SetUpdatedAt(t time.Time) *RouteControlUpdate {
	rcu.mutation.SetUpdatedAt(t)
	return rcu
}

// SetDistanceMethod sets the "distance_method" field.
func (rcu *RouteControlUpdate) SetDistanceMethod(rm routecontrol.DistanceMethod) *RouteControlUpdate {
	rcu.mutation.SetDistanceMethod(rm)
	return rcu
}

// SetNillableDistanceMethod sets the "distance_method" field if the given value is not nil.
func (rcu *RouteControlUpdate) SetNillableDistanceMethod(rm *routecontrol.DistanceMethod) *RouteControlUpdate {
	if rm != nil {
		rcu.SetDistanceMethod(*rm)
	}
	return rcu
}

// SetMileageUnit sets the "mileage_unit" field.
func (rcu *RouteControlUpdate) SetMileageUnit(ru routecontrol.MileageUnit) *RouteControlUpdate {
	rcu.mutation.SetMileageUnit(ru)
	return rcu
}

// SetNillableMileageUnit sets the "mileage_unit" field if the given value is not nil.
func (rcu *RouteControlUpdate) SetNillableMileageUnit(ru *routecontrol.MileageUnit) *RouteControlUpdate {
	if ru != nil {
		rcu.SetMileageUnit(*ru)
	}
	return rcu
}

// SetGenerateRoutes sets the "generate_routes" field.
func (rcu *RouteControlUpdate) SetGenerateRoutes(b bool) *RouteControlUpdate {
	rcu.mutation.SetGenerateRoutes(b)
	return rcu
}

// SetNillableGenerateRoutes sets the "generate_routes" field if the given value is not nil.
func (rcu *RouteControlUpdate) SetNillableGenerateRoutes(b *bool) *RouteControlUpdate {
	if b != nil {
		rcu.SetGenerateRoutes(*b)
	}
	return rcu
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (rcu *RouteControlUpdate) SetOrganizationID(id uuid.UUID) *RouteControlUpdate {
	rcu.mutation.SetOrganizationID(id)
	return rcu
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (rcu *RouteControlUpdate) SetOrganization(o *Organization) *RouteControlUpdate {
	return rcu.SetOrganizationID(o.ID)
}

// SetBusinessUnitID sets the "business_unit" edge to the BusinessUnit entity by ID.
func (rcu *RouteControlUpdate) SetBusinessUnitID(id uuid.UUID) *RouteControlUpdate {
	rcu.mutation.SetBusinessUnitID(id)
	return rcu
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (rcu *RouteControlUpdate) SetBusinessUnit(b *BusinessUnit) *RouteControlUpdate {
	return rcu.SetBusinessUnitID(b.ID)
}

// Mutation returns the RouteControlMutation object of the builder.
func (rcu *RouteControlUpdate) Mutation() *RouteControlMutation {
	return rcu.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (rcu *RouteControlUpdate) ClearOrganization() *RouteControlUpdate {
	rcu.mutation.ClearOrganization()
	return rcu
}

// ClearBusinessUnit clears the "business_unit" edge to the BusinessUnit entity.
func (rcu *RouteControlUpdate) ClearBusinessUnit() *RouteControlUpdate {
	rcu.mutation.ClearBusinessUnit()
	return rcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rcu *RouteControlUpdate) Save(ctx context.Context) (int, error) {
	rcu.defaults()
	return withHooks(ctx, rcu.sqlSave, rcu.mutation, rcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rcu *RouteControlUpdate) SaveX(ctx context.Context) int {
	affected, err := rcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rcu *RouteControlUpdate) Exec(ctx context.Context) error {
	_, err := rcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcu *RouteControlUpdate) ExecX(ctx context.Context) {
	if err := rcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rcu *RouteControlUpdate) defaults() {
	if _, ok := rcu.mutation.UpdatedAt(); !ok {
		v := routecontrol.UpdateDefaultUpdatedAt()
		rcu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rcu *RouteControlUpdate) check() error {
	if v, ok := rcu.mutation.DistanceMethod(); ok {
		if err := routecontrol.DistanceMethodValidator(v); err != nil {
			return &ValidationError{Name: "distance_method", err: fmt.Errorf(`ent: validator failed for field "RouteControl.distance_method": %w`, err)}
		}
	}
	if v, ok := rcu.mutation.MileageUnit(); ok {
		if err := routecontrol.MileageUnitValidator(v); err != nil {
			return &ValidationError{Name: "mileage_unit", err: fmt.Errorf(`ent: validator failed for field "RouteControl.mileage_unit": %w`, err)}
		}
	}
	if _, ok := rcu.mutation.OrganizationID(); rcu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RouteControl.organization"`)
	}
	if _, ok := rcu.mutation.BusinessUnitID(); rcu.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RouteControl.business_unit"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (rcu *RouteControlUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RouteControlUpdate {
	rcu.modifiers = append(rcu.modifiers, modifiers...)
	return rcu
}

func (rcu *RouteControlUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(routecontrol.Table, routecontrol.Columns, sqlgraph.NewFieldSpec(routecontrol.FieldID, field.TypeUUID))
	if ps := rcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rcu.mutation.UpdatedAt(); ok {
		_spec.SetField(routecontrol.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rcu.mutation.DistanceMethod(); ok {
		_spec.SetField(routecontrol.FieldDistanceMethod, field.TypeEnum, value)
	}
	if value, ok := rcu.mutation.MileageUnit(); ok {
		_spec.SetField(routecontrol.FieldMileageUnit, field.TypeEnum, value)
	}
	if value, ok := rcu.mutation.GenerateRoutes(); ok {
		_spec.SetField(routecontrol.FieldGenerateRoutes, field.TypeBool, value)
	}
	if rcu.mutation.OrganizationCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rcu.mutation.OrganizationIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rcu.mutation.BusinessUnitCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rcu.mutation.BusinessUnitIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(rcu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, rcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{routecontrol.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rcu.mutation.done = true
	return n, nil
}

// RouteControlUpdateOne is the builder for updating a single RouteControl entity.
type RouteControlUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *RouteControlMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (rcuo *RouteControlUpdateOne) SetUpdatedAt(t time.Time) *RouteControlUpdateOne {
	rcuo.mutation.SetUpdatedAt(t)
	return rcuo
}

// SetDistanceMethod sets the "distance_method" field.
func (rcuo *RouteControlUpdateOne) SetDistanceMethod(rm routecontrol.DistanceMethod) *RouteControlUpdateOne {
	rcuo.mutation.SetDistanceMethod(rm)
	return rcuo
}

// SetNillableDistanceMethod sets the "distance_method" field if the given value is not nil.
func (rcuo *RouteControlUpdateOne) SetNillableDistanceMethod(rm *routecontrol.DistanceMethod) *RouteControlUpdateOne {
	if rm != nil {
		rcuo.SetDistanceMethod(*rm)
	}
	return rcuo
}

// SetMileageUnit sets the "mileage_unit" field.
func (rcuo *RouteControlUpdateOne) SetMileageUnit(ru routecontrol.MileageUnit) *RouteControlUpdateOne {
	rcuo.mutation.SetMileageUnit(ru)
	return rcuo
}

// SetNillableMileageUnit sets the "mileage_unit" field if the given value is not nil.
func (rcuo *RouteControlUpdateOne) SetNillableMileageUnit(ru *routecontrol.MileageUnit) *RouteControlUpdateOne {
	if ru != nil {
		rcuo.SetMileageUnit(*ru)
	}
	return rcuo
}

// SetGenerateRoutes sets the "generate_routes" field.
func (rcuo *RouteControlUpdateOne) SetGenerateRoutes(b bool) *RouteControlUpdateOne {
	rcuo.mutation.SetGenerateRoutes(b)
	return rcuo
}

// SetNillableGenerateRoutes sets the "generate_routes" field if the given value is not nil.
func (rcuo *RouteControlUpdateOne) SetNillableGenerateRoutes(b *bool) *RouteControlUpdateOne {
	if b != nil {
		rcuo.SetGenerateRoutes(*b)
	}
	return rcuo
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (rcuo *RouteControlUpdateOne) SetOrganizationID(id uuid.UUID) *RouteControlUpdateOne {
	rcuo.mutation.SetOrganizationID(id)
	return rcuo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (rcuo *RouteControlUpdateOne) SetOrganization(o *Organization) *RouteControlUpdateOne {
	return rcuo.SetOrganizationID(o.ID)
}

// SetBusinessUnitID sets the "business_unit" edge to the BusinessUnit entity by ID.
func (rcuo *RouteControlUpdateOne) SetBusinessUnitID(id uuid.UUID) *RouteControlUpdateOne {
	rcuo.mutation.SetBusinessUnitID(id)
	return rcuo
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (rcuo *RouteControlUpdateOne) SetBusinessUnit(b *BusinessUnit) *RouteControlUpdateOne {
	return rcuo.SetBusinessUnitID(b.ID)
}

// Mutation returns the RouteControlMutation object of the builder.
func (rcuo *RouteControlUpdateOne) Mutation() *RouteControlMutation {
	return rcuo.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (rcuo *RouteControlUpdateOne) ClearOrganization() *RouteControlUpdateOne {
	rcuo.mutation.ClearOrganization()
	return rcuo
}

// ClearBusinessUnit clears the "business_unit" edge to the BusinessUnit entity.
func (rcuo *RouteControlUpdateOne) ClearBusinessUnit() *RouteControlUpdateOne {
	rcuo.mutation.ClearBusinessUnit()
	return rcuo
}

// Where appends a list predicates to the RouteControlUpdate builder.
func (rcuo *RouteControlUpdateOne) Where(ps ...predicate.RouteControl) *RouteControlUpdateOne {
	rcuo.mutation.Where(ps...)
	return rcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rcuo *RouteControlUpdateOne) Select(field string, fields ...string) *RouteControlUpdateOne {
	rcuo.fields = append([]string{field}, fields...)
	return rcuo
}

// Save executes the query and returns the updated RouteControl entity.
func (rcuo *RouteControlUpdateOne) Save(ctx context.Context) (*RouteControl, error) {
	rcuo.defaults()
	return withHooks(ctx, rcuo.sqlSave, rcuo.mutation, rcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rcuo *RouteControlUpdateOne) SaveX(ctx context.Context) *RouteControl {
	node, err := rcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rcuo *RouteControlUpdateOne) Exec(ctx context.Context) error {
	_, err := rcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcuo *RouteControlUpdateOne) ExecX(ctx context.Context) {
	if err := rcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rcuo *RouteControlUpdateOne) defaults() {
	if _, ok := rcuo.mutation.UpdatedAt(); !ok {
		v := routecontrol.UpdateDefaultUpdatedAt()
		rcuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rcuo *RouteControlUpdateOne) check() error {
	if v, ok := rcuo.mutation.DistanceMethod(); ok {
		if err := routecontrol.DistanceMethodValidator(v); err != nil {
			return &ValidationError{Name: "distance_method", err: fmt.Errorf(`ent: validator failed for field "RouteControl.distance_method": %w`, err)}
		}
	}
	if v, ok := rcuo.mutation.MileageUnit(); ok {
		if err := routecontrol.MileageUnitValidator(v); err != nil {
			return &ValidationError{Name: "mileage_unit", err: fmt.Errorf(`ent: validator failed for field "RouteControl.mileage_unit": %w`, err)}
		}
	}
	if _, ok := rcuo.mutation.OrganizationID(); rcuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RouteControl.organization"`)
	}
	if _, ok := rcuo.mutation.BusinessUnitID(); rcuo.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RouteControl.business_unit"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (rcuo *RouteControlUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RouteControlUpdateOne {
	rcuo.modifiers = append(rcuo.modifiers, modifiers...)
	return rcuo
}

func (rcuo *RouteControlUpdateOne) sqlSave(ctx context.Context) (_node *RouteControl, err error) {
	if err := rcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(routecontrol.Table, routecontrol.Columns, sqlgraph.NewFieldSpec(routecontrol.FieldID, field.TypeUUID))
	id, ok := rcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RouteControl.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, routecontrol.FieldID)
		for _, f := range fields {
			if !routecontrol.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != routecontrol.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(routecontrol.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rcuo.mutation.DistanceMethod(); ok {
		_spec.SetField(routecontrol.FieldDistanceMethod, field.TypeEnum, value)
	}
	if value, ok := rcuo.mutation.MileageUnit(); ok {
		_spec.SetField(routecontrol.FieldMileageUnit, field.TypeEnum, value)
	}
	if value, ok := rcuo.mutation.GenerateRoutes(); ok {
		_spec.SetField(routecontrol.FieldGenerateRoutes, field.TypeBool, value)
	}
	if rcuo.mutation.OrganizationCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rcuo.mutation.OrganizationIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rcuo.mutation.BusinessUnitCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rcuo.mutation.BusinessUnitIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(rcuo.modifiers...)
	_node = &RouteControl{config: rcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{routecontrol.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rcuo.mutation.done = true
	return _node, nil
}
