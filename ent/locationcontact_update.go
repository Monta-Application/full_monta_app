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
	"github.com/emoss08/trenova/ent/location"
	"github.com/emoss08/trenova/ent/locationcontact"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// LocationContactUpdate is the builder for updating LocationContact entities.
type LocationContactUpdate struct {
	config
	hooks     []Hook
	mutation  *LocationContactMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the LocationContactUpdate builder.
func (lcu *LocationContactUpdate) Where(ps ...predicate.LocationContact) *LocationContactUpdate {
	lcu.mutation.Where(ps...)
	return lcu
}

// SetUpdatedAt sets the "updated_at" field.
func (lcu *LocationContactUpdate) SetUpdatedAt(t time.Time) *LocationContactUpdate {
	lcu.mutation.SetUpdatedAt(t)
	return lcu
}

// SetVersion sets the "version" field.
func (lcu *LocationContactUpdate) SetVersion(i int) *LocationContactUpdate {
	lcu.mutation.ResetVersion()
	lcu.mutation.SetVersion(i)
	return lcu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (lcu *LocationContactUpdate) SetNillableVersion(i *int) *LocationContactUpdate {
	if i != nil {
		lcu.SetVersion(*i)
	}
	return lcu
}

// AddVersion adds i to the "version" field.
func (lcu *LocationContactUpdate) AddVersion(i int) *LocationContactUpdate {
	lcu.mutation.AddVersion(i)
	return lcu
}

// SetLocationID sets the "location_id" field.
func (lcu *LocationContactUpdate) SetLocationID(u uuid.UUID) *LocationContactUpdate {
	lcu.mutation.SetLocationID(u)
	return lcu
}

// SetNillableLocationID sets the "location_id" field if the given value is not nil.
func (lcu *LocationContactUpdate) SetNillableLocationID(u *uuid.UUID) *LocationContactUpdate {
	if u != nil {
		lcu.SetLocationID(*u)
	}
	return lcu
}

// SetName sets the "name" field.
func (lcu *LocationContactUpdate) SetName(s string) *LocationContactUpdate {
	lcu.mutation.SetName(s)
	return lcu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (lcu *LocationContactUpdate) SetNillableName(s *string) *LocationContactUpdate {
	if s != nil {
		lcu.SetName(*s)
	}
	return lcu
}

// SetEmailAddress sets the "email_address" field.
func (lcu *LocationContactUpdate) SetEmailAddress(s string) *LocationContactUpdate {
	lcu.mutation.SetEmailAddress(s)
	return lcu
}

// SetNillableEmailAddress sets the "email_address" field if the given value is not nil.
func (lcu *LocationContactUpdate) SetNillableEmailAddress(s *string) *LocationContactUpdate {
	if s != nil {
		lcu.SetEmailAddress(*s)
	}
	return lcu
}

// ClearEmailAddress clears the value of the "email_address" field.
func (lcu *LocationContactUpdate) ClearEmailAddress() *LocationContactUpdate {
	lcu.mutation.ClearEmailAddress()
	return lcu
}

// SetPhoneNumber sets the "phone_number" field.
func (lcu *LocationContactUpdate) SetPhoneNumber(s string) *LocationContactUpdate {
	lcu.mutation.SetPhoneNumber(s)
	return lcu
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (lcu *LocationContactUpdate) SetNillablePhoneNumber(s *string) *LocationContactUpdate {
	if s != nil {
		lcu.SetPhoneNumber(*s)
	}
	return lcu
}

// ClearPhoneNumber clears the value of the "phone_number" field.
func (lcu *LocationContactUpdate) ClearPhoneNumber() *LocationContactUpdate {
	lcu.mutation.ClearPhoneNumber()
	return lcu
}

// SetLocation sets the "location" edge to the Location entity.
func (lcu *LocationContactUpdate) SetLocation(l *Location) *LocationContactUpdate {
	return lcu.SetLocationID(l.ID)
}

// Mutation returns the LocationContactMutation object of the builder.
func (lcu *LocationContactUpdate) Mutation() *LocationContactMutation {
	return lcu.mutation
}

// ClearLocation clears the "location" edge to the Location entity.
func (lcu *LocationContactUpdate) ClearLocation() *LocationContactUpdate {
	lcu.mutation.ClearLocation()
	return lcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lcu *LocationContactUpdate) Save(ctx context.Context) (int, error) {
	lcu.defaults()
	return withHooks(ctx, lcu.sqlSave, lcu.mutation, lcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lcu *LocationContactUpdate) SaveX(ctx context.Context) int {
	affected, err := lcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lcu *LocationContactUpdate) Exec(ctx context.Context) error {
	_, err := lcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcu *LocationContactUpdate) ExecX(ctx context.Context) {
	if err := lcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lcu *LocationContactUpdate) defaults() {
	if _, ok := lcu.mutation.UpdatedAt(); !ok {
		v := locationcontact.UpdateDefaultUpdatedAt()
		lcu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lcu *LocationContactUpdate) check() error {
	if v, ok := lcu.mutation.Name(); ok {
		if err := locationcontact.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "LocationContact.name": %w`, err)}
		}
	}
	if v, ok := lcu.mutation.PhoneNumber(); ok {
		if err := locationcontact.PhoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "phone_number", err: fmt.Errorf(`ent: validator failed for field "LocationContact.phone_number": %w`, err)}
		}
	}
	if _, ok := lcu.mutation.BusinessUnitID(); lcu.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LocationContact.business_unit"`)
	}
	if _, ok := lcu.mutation.OrganizationID(); lcu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LocationContact.organization"`)
	}
	if _, ok := lcu.mutation.LocationID(); lcu.mutation.LocationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LocationContact.location"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (lcu *LocationContactUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LocationContactUpdate {
	lcu.modifiers = append(lcu.modifiers, modifiers...)
	return lcu
}

func (lcu *LocationContactUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(locationcontact.Table, locationcontact.Columns, sqlgraph.NewFieldSpec(locationcontact.FieldID, field.TypeUUID))
	if ps := lcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lcu.mutation.UpdatedAt(); ok {
		_spec.SetField(locationcontact.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lcu.mutation.Version(); ok {
		_spec.SetField(locationcontact.FieldVersion, field.TypeInt, value)
	}
	if value, ok := lcu.mutation.AddedVersion(); ok {
		_spec.AddField(locationcontact.FieldVersion, field.TypeInt, value)
	}
	if value, ok := lcu.mutation.Name(); ok {
		_spec.SetField(locationcontact.FieldName, field.TypeString, value)
	}
	if value, ok := lcu.mutation.EmailAddress(); ok {
		_spec.SetField(locationcontact.FieldEmailAddress, field.TypeString, value)
	}
	if lcu.mutation.EmailAddressCleared() {
		_spec.ClearField(locationcontact.FieldEmailAddress, field.TypeString)
	}
	if value, ok := lcu.mutation.PhoneNumber(); ok {
		_spec.SetField(locationcontact.FieldPhoneNumber, field.TypeString, value)
	}
	if lcu.mutation.PhoneNumberCleared() {
		_spec.ClearField(locationcontact.FieldPhoneNumber, field.TypeString)
	}
	if lcu.mutation.LocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   locationcontact.LocationTable,
			Columns: []string{locationcontact.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lcu.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   locationcontact.LocationTable,
			Columns: []string{locationcontact.LocationColumn},
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
	_spec.AddModifiers(lcu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, lcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{locationcontact.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lcu.mutation.done = true
	return n, nil
}

// LocationContactUpdateOne is the builder for updating a single LocationContact entity.
type LocationContactUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *LocationContactMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (lcuo *LocationContactUpdateOne) SetUpdatedAt(t time.Time) *LocationContactUpdateOne {
	lcuo.mutation.SetUpdatedAt(t)
	return lcuo
}

// SetVersion sets the "version" field.
func (lcuo *LocationContactUpdateOne) SetVersion(i int) *LocationContactUpdateOne {
	lcuo.mutation.ResetVersion()
	lcuo.mutation.SetVersion(i)
	return lcuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (lcuo *LocationContactUpdateOne) SetNillableVersion(i *int) *LocationContactUpdateOne {
	if i != nil {
		lcuo.SetVersion(*i)
	}
	return lcuo
}

// AddVersion adds i to the "version" field.
func (lcuo *LocationContactUpdateOne) AddVersion(i int) *LocationContactUpdateOne {
	lcuo.mutation.AddVersion(i)
	return lcuo
}

// SetLocationID sets the "location_id" field.
func (lcuo *LocationContactUpdateOne) SetLocationID(u uuid.UUID) *LocationContactUpdateOne {
	lcuo.mutation.SetLocationID(u)
	return lcuo
}

// SetNillableLocationID sets the "location_id" field if the given value is not nil.
func (lcuo *LocationContactUpdateOne) SetNillableLocationID(u *uuid.UUID) *LocationContactUpdateOne {
	if u != nil {
		lcuo.SetLocationID(*u)
	}
	return lcuo
}

// SetName sets the "name" field.
func (lcuo *LocationContactUpdateOne) SetName(s string) *LocationContactUpdateOne {
	lcuo.mutation.SetName(s)
	return lcuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (lcuo *LocationContactUpdateOne) SetNillableName(s *string) *LocationContactUpdateOne {
	if s != nil {
		lcuo.SetName(*s)
	}
	return lcuo
}

// SetEmailAddress sets the "email_address" field.
func (lcuo *LocationContactUpdateOne) SetEmailAddress(s string) *LocationContactUpdateOne {
	lcuo.mutation.SetEmailAddress(s)
	return lcuo
}

// SetNillableEmailAddress sets the "email_address" field if the given value is not nil.
func (lcuo *LocationContactUpdateOne) SetNillableEmailAddress(s *string) *LocationContactUpdateOne {
	if s != nil {
		lcuo.SetEmailAddress(*s)
	}
	return lcuo
}

// ClearEmailAddress clears the value of the "email_address" field.
func (lcuo *LocationContactUpdateOne) ClearEmailAddress() *LocationContactUpdateOne {
	lcuo.mutation.ClearEmailAddress()
	return lcuo
}

// SetPhoneNumber sets the "phone_number" field.
func (lcuo *LocationContactUpdateOne) SetPhoneNumber(s string) *LocationContactUpdateOne {
	lcuo.mutation.SetPhoneNumber(s)
	return lcuo
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (lcuo *LocationContactUpdateOne) SetNillablePhoneNumber(s *string) *LocationContactUpdateOne {
	if s != nil {
		lcuo.SetPhoneNumber(*s)
	}
	return lcuo
}

// ClearPhoneNumber clears the value of the "phone_number" field.
func (lcuo *LocationContactUpdateOne) ClearPhoneNumber() *LocationContactUpdateOne {
	lcuo.mutation.ClearPhoneNumber()
	return lcuo
}

// SetLocation sets the "location" edge to the Location entity.
func (lcuo *LocationContactUpdateOne) SetLocation(l *Location) *LocationContactUpdateOne {
	return lcuo.SetLocationID(l.ID)
}

// Mutation returns the LocationContactMutation object of the builder.
func (lcuo *LocationContactUpdateOne) Mutation() *LocationContactMutation {
	return lcuo.mutation
}

// ClearLocation clears the "location" edge to the Location entity.
func (lcuo *LocationContactUpdateOne) ClearLocation() *LocationContactUpdateOne {
	lcuo.mutation.ClearLocation()
	return lcuo
}

// Where appends a list predicates to the LocationContactUpdate builder.
func (lcuo *LocationContactUpdateOne) Where(ps ...predicate.LocationContact) *LocationContactUpdateOne {
	lcuo.mutation.Where(ps...)
	return lcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lcuo *LocationContactUpdateOne) Select(field string, fields ...string) *LocationContactUpdateOne {
	lcuo.fields = append([]string{field}, fields...)
	return lcuo
}

// Save executes the query and returns the updated LocationContact entity.
func (lcuo *LocationContactUpdateOne) Save(ctx context.Context) (*LocationContact, error) {
	lcuo.defaults()
	return withHooks(ctx, lcuo.sqlSave, lcuo.mutation, lcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lcuo *LocationContactUpdateOne) SaveX(ctx context.Context) *LocationContact {
	node, err := lcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lcuo *LocationContactUpdateOne) Exec(ctx context.Context) error {
	_, err := lcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcuo *LocationContactUpdateOne) ExecX(ctx context.Context) {
	if err := lcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lcuo *LocationContactUpdateOne) defaults() {
	if _, ok := lcuo.mutation.UpdatedAt(); !ok {
		v := locationcontact.UpdateDefaultUpdatedAt()
		lcuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lcuo *LocationContactUpdateOne) check() error {
	if v, ok := lcuo.mutation.Name(); ok {
		if err := locationcontact.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "LocationContact.name": %w`, err)}
		}
	}
	if v, ok := lcuo.mutation.PhoneNumber(); ok {
		if err := locationcontact.PhoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "phone_number", err: fmt.Errorf(`ent: validator failed for field "LocationContact.phone_number": %w`, err)}
		}
	}
	if _, ok := lcuo.mutation.BusinessUnitID(); lcuo.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LocationContact.business_unit"`)
	}
	if _, ok := lcuo.mutation.OrganizationID(); lcuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LocationContact.organization"`)
	}
	if _, ok := lcuo.mutation.LocationID(); lcuo.mutation.LocationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LocationContact.location"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (lcuo *LocationContactUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LocationContactUpdateOne {
	lcuo.modifiers = append(lcuo.modifiers, modifiers...)
	return lcuo
}

func (lcuo *LocationContactUpdateOne) sqlSave(ctx context.Context) (_node *LocationContact, err error) {
	if err := lcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(locationcontact.Table, locationcontact.Columns, sqlgraph.NewFieldSpec(locationcontact.FieldID, field.TypeUUID))
	id, ok := lcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LocationContact.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, locationcontact.FieldID)
		for _, f := range fields {
			if !locationcontact.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != locationcontact.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(locationcontact.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lcuo.mutation.Version(); ok {
		_spec.SetField(locationcontact.FieldVersion, field.TypeInt, value)
	}
	if value, ok := lcuo.mutation.AddedVersion(); ok {
		_spec.AddField(locationcontact.FieldVersion, field.TypeInt, value)
	}
	if value, ok := lcuo.mutation.Name(); ok {
		_spec.SetField(locationcontact.FieldName, field.TypeString, value)
	}
	if value, ok := lcuo.mutation.EmailAddress(); ok {
		_spec.SetField(locationcontact.FieldEmailAddress, field.TypeString, value)
	}
	if lcuo.mutation.EmailAddressCleared() {
		_spec.ClearField(locationcontact.FieldEmailAddress, field.TypeString)
	}
	if value, ok := lcuo.mutation.PhoneNumber(); ok {
		_spec.SetField(locationcontact.FieldPhoneNumber, field.TypeString, value)
	}
	if lcuo.mutation.PhoneNumberCleared() {
		_spec.ClearField(locationcontact.FieldPhoneNumber, field.TypeString)
	}
	if lcuo.mutation.LocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   locationcontact.LocationTable,
			Columns: []string{locationcontact.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lcuo.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   locationcontact.LocationTable,
			Columns: []string{locationcontact.LocationColumn},
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
	_spec.AddModifiers(lcuo.modifiers...)
	_node = &LocationContact{config: lcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{locationcontact.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	lcuo.mutation.done = true
	return _node, nil
}
