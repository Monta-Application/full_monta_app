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
	"github.com/emoss08/trenova/internal/ent/deliveryslot"
	"github.com/emoss08/trenova/internal/ent/location"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/emoss08/trenova/internal/util/types"
	"github.com/google/uuid"
)

// DeliverySlotUpdate is the builder for updating DeliverySlot entities.
type DeliverySlotUpdate struct {
	config
	hooks     []Hook
	mutation  *DeliverySlotMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the DeliverySlotUpdate builder.
func (dsu *DeliverySlotUpdate) Where(ps ...predicate.DeliverySlot) *DeliverySlotUpdate {
	dsu.mutation.Where(ps...)
	return dsu
}

// SetUpdatedAt sets the "updated_at" field.
func (dsu *DeliverySlotUpdate) SetUpdatedAt(t time.Time) *DeliverySlotUpdate {
	dsu.mutation.SetUpdatedAt(t)
	return dsu
}

// SetVersion sets the "version" field.
func (dsu *DeliverySlotUpdate) SetVersion(i int) *DeliverySlotUpdate {
	dsu.mutation.ResetVersion()
	dsu.mutation.SetVersion(i)
	return dsu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (dsu *DeliverySlotUpdate) SetNillableVersion(i *int) *DeliverySlotUpdate {
	if i != nil {
		dsu.SetVersion(*i)
	}
	return dsu
}

// AddVersion adds i to the "version" field.
func (dsu *DeliverySlotUpdate) AddVersion(i int) *DeliverySlotUpdate {
	dsu.mutation.AddVersion(i)
	return dsu
}

// SetLocationID sets the "location_id" field.
func (dsu *DeliverySlotUpdate) SetLocationID(u uuid.UUID) *DeliverySlotUpdate {
	dsu.mutation.SetLocationID(u)
	return dsu
}

// SetNillableLocationID sets the "location_id" field if the given value is not nil.
func (dsu *DeliverySlotUpdate) SetNillableLocationID(u *uuid.UUID) *DeliverySlotUpdate {
	if u != nil {
		dsu.SetLocationID(*u)
	}
	return dsu
}

// SetDayOfWeek sets the "day_of_week" field.
func (dsu *DeliverySlotUpdate) SetDayOfWeek(dow deliveryslot.DayOfWeek) *DeliverySlotUpdate {
	dsu.mutation.SetDayOfWeek(dow)
	return dsu
}

// SetNillableDayOfWeek sets the "day_of_week" field if the given value is not nil.
func (dsu *DeliverySlotUpdate) SetNillableDayOfWeek(dow *deliveryslot.DayOfWeek) *DeliverySlotUpdate {
	if dow != nil {
		dsu.SetDayOfWeek(*dow)
	}
	return dsu
}

// SetStartTime sets the "start_time" field.
func (dsu *DeliverySlotUpdate) SetStartTime(to *types.TimeOnly) *DeliverySlotUpdate {
	dsu.mutation.SetStartTime(to)
	return dsu
}

// SetEndTime sets the "end_time" field.
func (dsu *DeliverySlotUpdate) SetEndTime(to *types.TimeOnly) *DeliverySlotUpdate {
	dsu.mutation.SetEndTime(to)
	return dsu
}

// SetLocation sets the "location" edge to the Location entity.
func (dsu *DeliverySlotUpdate) SetLocation(l *Location) *DeliverySlotUpdate {
	return dsu.SetLocationID(l.ID)
}

// Mutation returns the DeliverySlotMutation object of the builder.
func (dsu *DeliverySlotUpdate) Mutation() *DeliverySlotMutation {
	return dsu.mutation
}

// ClearLocation clears the "location" edge to the Location entity.
func (dsu *DeliverySlotUpdate) ClearLocation() *DeliverySlotUpdate {
	dsu.mutation.ClearLocation()
	return dsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dsu *DeliverySlotUpdate) Save(ctx context.Context) (int, error) {
	dsu.defaults()
	return withHooks(ctx, dsu.sqlSave, dsu.mutation, dsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dsu *DeliverySlotUpdate) SaveX(ctx context.Context) int {
	affected, err := dsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dsu *DeliverySlotUpdate) Exec(ctx context.Context) error {
	_, err := dsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dsu *DeliverySlotUpdate) ExecX(ctx context.Context) {
	if err := dsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dsu *DeliverySlotUpdate) defaults() {
	if _, ok := dsu.mutation.UpdatedAt(); !ok {
		v := deliveryslot.UpdateDefaultUpdatedAt()
		dsu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dsu *DeliverySlotUpdate) check() error {
	if v, ok := dsu.mutation.DayOfWeek(); ok {
		if err := deliveryslot.DayOfWeekValidator(v); err != nil {
			return &ValidationError{Name: "day_of_week", err: fmt.Errorf(`ent: validator failed for field "DeliverySlot.day_of_week": %w`, err)}
		}
	}
	if _, ok := dsu.mutation.BusinessUnitID(); dsu.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeliverySlot.business_unit"`)
	}
	if _, ok := dsu.mutation.OrganizationID(); dsu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeliverySlot.organization"`)
	}
	if _, ok := dsu.mutation.CustomerID(); dsu.mutation.CustomerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeliverySlot.customer"`)
	}
	if _, ok := dsu.mutation.LocationID(); dsu.mutation.LocationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeliverySlot.location"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (dsu *DeliverySlotUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DeliverySlotUpdate {
	dsu.modifiers = append(dsu.modifiers, modifiers...)
	return dsu
}

func (dsu *DeliverySlotUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := dsu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(deliveryslot.Table, deliveryslot.Columns, sqlgraph.NewFieldSpec(deliveryslot.FieldID, field.TypeUUID))
	if ps := dsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dsu.mutation.UpdatedAt(); ok {
		_spec.SetField(deliveryslot.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := dsu.mutation.Version(); ok {
		_spec.SetField(deliveryslot.FieldVersion, field.TypeInt, value)
	}
	if value, ok := dsu.mutation.AddedVersion(); ok {
		_spec.AddField(deliveryslot.FieldVersion, field.TypeInt, value)
	}
	if value, ok := dsu.mutation.DayOfWeek(); ok {
		_spec.SetField(deliveryslot.FieldDayOfWeek, field.TypeEnum, value)
	}
	if value, ok := dsu.mutation.StartTime(); ok {
		_spec.SetField(deliveryslot.FieldStartTime, field.TypeOther, value)
	}
	if value, ok := dsu.mutation.EndTime(); ok {
		_spec.SetField(deliveryslot.FieldEndTime, field.TypeOther, value)
	}
	if dsu.mutation.LocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deliveryslot.LocationTable,
			Columns: []string{deliveryslot.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dsu.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deliveryslot.LocationTable,
			Columns: []string{deliveryslot.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(dsu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, dsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deliveryslot.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	dsu.mutation.done = true
	return n, nil
}

// DeliverySlotUpdateOne is the builder for updating a single DeliverySlot entity.
type DeliverySlotUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *DeliverySlotMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (dsuo *DeliverySlotUpdateOne) SetUpdatedAt(t time.Time) *DeliverySlotUpdateOne {
	dsuo.mutation.SetUpdatedAt(t)
	return dsuo
}

// SetVersion sets the "version" field.
func (dsuo *DeliverySlotUpdateOne) SetVersion(i int) *DeliverySlotUpdateOne {
	dsuo.mutation.ResetVersion()
	dsuo.mutation.SetVersion(i)
	return dsuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (dsuo *DeliverySlotUpdateOne) SetNillableVersion(i *int) *DeliverySlotUpdateOne {
	if i != nil {
		dsuo.SetVersion(*i)
	}
	return dsuo
}

// AddVersion adds i to the "version" field.
func (dsuo *DeliverySlotUpdateOne) AddVersion(i int) *DeliverySlotUpdateOne {
	dsuo.mutation.AddVersion(i)
	return dsuo
}

// SetLocationID sets the "location_id" field.
func (dsuo *DeliverySlotUpdateOne) SetLocationID(u uuid.UUID) *DeliverySlotUpdateOne {
	dsuo.mutation.SetLocationID(u)
	return dsuo
}

// SetNillableLocationID sets the "location_id" field if the given value is not nil.
func (dsuo *DeliverySlotUpdateOne) SetNillableLocationID(u *uuid.UUID) *DeliverySlotUpdateOne {
	if u != nil {
		dsuo.SetLocationID(*u)
	}
	return dsuo
}

// SetDayOfWeek sets the "day_of_week" field.
func (dsuo *DeliverySlotUpdateOne) SetDayOfWeek(dow deliveryslot.DayOfWeek) *DeliverySlotUpdateOne {
	dsuo.mutation.SetDayOfWeek(dow)
	return dsuo
}

// SetNillableDayOfWeek sets the "day_of_week" field if the given value is not nil.
func (dsuo *DeliverySlotUpdateOne) SetNillableDayOfWeek(dow *deliveryslot.DayOfWeek) *DeliverySlotUpdateOne {
	if dow != nil {
		dsuo.SetDayOfWeek(*dow)
	}
	return dsuo
}

// SetStartTime sets the "start_time" field.
func (dsuo *DeliverySlotUpdateOne) SetStartTime(to *types.TimeOnly) *DeliverySlotUpdateOne {
	dsuo.mutation.SetStartTime(to)
	return dsuo
}

// SetEndTime sets the "end_time" field.
func (dsuo *DeliverySlotUpdateOne) SetEndTime(to *types.TimeOnly) *DeliverySlotUpdateOne {
	dsuo.mutation.SetEndTime(to)
	return dsuo
}

// SetLocation sets the "location" edge to the Location entity.
func (dsuo *DeliverySlotUpdateOne) SetLocation(l *Location) *DeliverySlotUpdateOne {
	return dsuo.SetLocationID(l.ID)
}

// Mutation returns the DeliverySlotMutation object of the builder.
func (dsuo *DeliverySlotUpdateOne) Mutation() *DeliverySlotMutation {
	return dsuo.mutation
}

// ClearLocation clears the "location" edge to the Location entity.
func (dsuo *DeliverySlotUpdateOne) ClearLocation() *DeliverySlotUpdateOne {
	dsuo.mutation.ClearLocation()
	return dsuo
}

// Where appends a list predicates to the DeliverySlotUpdate builder.
func (dsuo *DeliverySlotUpdateOne) Where(ps ...predicate.DeliverySlot) *DeliverySlotUpdateOne {
	dsuo.mutation.Where(ps...)
	return dsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dsuo *DeliverySlotUpdateOne) Select(field string, fields ...string) *DeliverySlotUpdateOne {
	dsuo.fields = append([]string{field}, fields...)
	return dsuo
}

// Save executes the query and returns the updated DeliverySlot entity.
func (dsuo *DeliverySlotUpdateOne) Save(ctx context.Context) (*DeliverySlot, error) {
	dsuo.defaults()
	return withHooks(ctx, dsuo.sqlSave, dsuo.mutation, dsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dsuo *DeliverySlotUpdateOne) SaveX(ctx context.Context) *DeliverySlot {
	node, err := dsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dsuo *DeliverySlotUpdateOne) Exec(ctx context.Context) error {
	_, err := dsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dsuo *DeliverySlotUpdateOne) ExecX(ctx context.Context) {
	if err := dsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dsuo *DeliverySlotUpdateOne) defaults() {
	if _, ok := dsuo.mutation.UpdatedAt(); !ok {
		v := deliveryslot.UpdateDefaultUpdatedAt()
		dsuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dsuo *DeliverySlotUpdateOne) check() error {
	if v, ok := dsuo.mutation.DayOfWeek(); ok {
		if err := deliveryslot.DayOfWeekValidator(v); err != nil {
			return &ValidationError{Name: "day_of_week", err: fmt.Errorf(`ent: validator failed for field "DeliverySlot.day_of_week": %w`, err)}
		}
	}
	if _, ok := dsuo.mutation.BusinessUnitID(); dsuo.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeliverySlot.business_unit"`)
	}
	if _, ok := dsuo.mutation.OrganizationID(); dsuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeliverySlot.organization"`)
	}
	if _, ok := dsuo.mutation.CustomerID(); dsuo.mutation.CustomerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeliverySlot.customer"`)
	}
	if _, ok := dsuo.mutation.LocationID(); dsuo.mutation.LocationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeliverySlot.location"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (dsuo *DeliverySlotUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DeliverySlotUpdateOne {
	dsuo.modifiers = append(dsuo.modifiers, modifiers...)
	return dsuo
}

func (dsuo *DeliverySlotUpdateOne) sqlSave(ctx context.Context) (_node *DeliverySlot, err error) {
	if err := dsuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(deliveryslot.Table, deliveryslot.Columns, sqlgraph.NewFieldSpec(deliveryslot.FieldID, field.TypeUUID))
	id, ok := dsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DeliverySlot.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deliveryslot.FieldID)
		for _, f := range fields {
			if !deliveryslot.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != deliveryslot.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dsuo.mutation.UpdatedAt(); ok {
		_spec.SetField(deliveryslot.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := dsuo.mutation.Version(); ok {
		_spec.SetField(deliveryslot.FieldVersion, field.TypeInt, value)
	}
	if value, ok := dsuo.mutation.AddedVersion(); ok {
		_spec.AddField(deliveryslot.FieldVersion, field.TypeInt, value)
	}
	if value, ok := dsuo.mutation.DayOfWeek(); ok {
		_spec.SetField(deliveryslot.FieldDayOfWeek, field.TypeEnum, value)
	}
	if value, ok := dsuo.mutation.StartTime(); ok {
		_spec.SetField(deliveryslot.FieldStartTime, field.TypeOther, value)
	}
	if value, ok := dsuo.mutation.EndTime(); ok {
		_spec.SetField(deliveryslot.FieldEndTime, field.TypeOther, value)
	}
	if dsuo.mutation.LocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deliveryslot.LocationTable,
			Columns: []string{deliveryslot.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dsuo.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deliveryslot.LocationTable,
			Columns: []string{deliveryslot.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(dsuo.modifiers...)
	_node = &DeliverySlot{config: dsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deliveryslot.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	dsuo.mutation.done = true
	return _node, nil
}
