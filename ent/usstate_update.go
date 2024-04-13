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
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/emoss08/trenova/ent/usstate"
)

// UsStateUpdate is the builder for updating UsState entities.
type UsStateUpdate struct {
	config
	hooks     []Hook
	mutation  *UsStateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the UsStateUpdate builder.
func (usu *UsStateUpdate) Where(ps ...predicate.UsState) *UsStateUpdate {
	usu.mutation.Where(ps...)
	return usu
}

// SetUpdatedAt sets the "updated_at" field.
func (usu *UsStateUpdate) SetUpdatedAt(t time.Time) *UsStateUpdate {
	usu.mutation.SetUpdatedAt(t)
	return usu
}

// SetName sets the "name" field.
func (usu *UsStateUpdate) SetName(s string) *UsStateUpdate {
	usu.mutation.SetName(s)
	return usu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (usu *UsStateUpdate) SetNillableName(s *string) *UsStateUpdate {
	if s != nil {
		usu.SetName(*s)
	}
	return usu
}

// SetAbbreviation sets the "abbreviation" field.
func (usu *UsStateUpdate) SetAbbreviation(s string) *UsStateUpdate {
	usu.mutation.SetAbbreviation(s)
	return usu
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (usu *UsStateUpdate) SetNillableAbbreviation(s *string) *UsStateUpdate {
	if s != nil {
		usu.SetAbbreviation(*s)
	}
	return usu
}

// SetCountryName sets the "country_name" field.
func (usu *UsStateUpdate) SetCountryName(s string) *UsStateUpdate {
	usu.mutation.SetCountryName(s)
	return usu
}

// SetNillableCountryName sets the "country_name" field if the given value is not nil.
func (usu *UsStateUpdate) SetNillableCountryName(s *string) *UsStateUpdate {
	if s != nil {
		usu.SetCountryName(*s)
	}
	return usu
}

// SetCountryIso3 sets the "country_iso3" field.
func (usu *UsStateUpdate) SetCountryIso3(s string) *UsStateUpdate {
	usu.mutation.SetCountryIso3(s)
	return usu
}

// SetNillableCountryIso3 sets the "country_iso3" field if the given value is not nil.
func (usu *UsStateUpdate) SetNillableCountryIso3(s *string) *UsStateUpdate {
	if s != nil {
		usu.SetCountryIso3(*s)
	}
	return usu
}

// Mutation returns the UsStateMutation object of the builder.
func (usu *UsStateUpdate) Mutation() *UsStateMutation {
	return usu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (usu *UsStateUpdate) Save(ctx context.Context) (int, error) {
	usu.defaults()
	return withHooks(ctx, usu.sqlSave, usu.mutation, usu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (usu *UsStateUpdate) SaveX(ctx context.Context) int {
	affected, err := usu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (usu *UsStateUpdate) Exec(ctx context.Context) error {
	_, err := usu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usu *UsStateUpdate) ExecX(ctx context.Context) {
	if err := usu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usu *UsStateUpdate) defaults() {
	if _, ok := usu.mutation.UpdatedAt(); !ok {
		v := usstate.UpdateDefaultUpdatedAt()
		usu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usu *UsStateUpdate) check() error {
	if v, ok := usu.mutation.Name(); ok {
		if err := usstate.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "UsState.name": %w`, err)}
		}
	}
	if v, ok := usu.mutation.Abbreviation(); ok {
		if err := usstate.AbbreviationValidator(v); err != nil {
			return &ValidationError{Name: "abbreviation", err: fmt.Errorf(`ent: validator failed for field "UsState.abbreviation": %w`, err)}
		}
	}
	if v, ok := usu.mutation.CountryName(); ok {
		if err := usstate.CountryNameValidator(v); err != nil {
			return &ValidationError{Name: "country_name", err: fmt.Errorf(`ent: validator failed for field "UsState.country_name": %w`, err)}
		}
	}
	if v, ok := usu.mutation.CountryIso3(); ok {
		if err := usstate.CountryIso3Validator(v); err != nil {
			return &ValidationError{Name: "country_iso3", err: fmt.Errorf(`ent: validator failed for field "UsState.country_iso3": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (usu *UsStateUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UsStateUpdate {
	usu.modifiers = append(usu.modifiers, modifiers...)
	return usu
}

func (usu *UsStateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := usu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(usstate.Table, usstate.Columns, sqlgraph.NewFieldSpec(usstate.FieldID, field.TypeUUID))
	if ps := usu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := usu.mutation.UpdatedAt(); ok {
		_spec.SetField(usstate.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := usu.mutation.Name(); ok {
		_spec.SetField(usstate.FieldName, field.TypeString, value)
	}
	if value, ok := usu.mutation.Abbreviation(); ok {
		_spec.SetField(usstate.FieldAbbreviation, field.TypeString, value)
	}
	if value, ok := usu.mutation.CountryName(); ok {
		_spec.SetField(usstate.FieldCountryName, field.TypeString, value)
	}
	if value, ok := usu.mutation.CountryIso3(); ok {
		_spec.SetField(usstate.FieldCountryIso3, field.TypeString, value)
	}
	_spec.AddModifiers(usu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, usu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usstate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	usu.mutation.done = true
	return n, nil
}

// UsStateUpdateOne is the builder for updating a single UsState entity.
type UsStateUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *UsStateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (usuo *UsStateUpdateOne) SetUpdatedAt(t time.Time) *UsStateUpdateOne {
	usuo.mutation.SetUpdatedAt(t)
	return usuo
}

// SetName sets the "name" field.
func (usuo *UsStateUpdateOne) SetName(s string) *UsStateUpdateOne {
	usuo.mutation.SetName(s)
	return usuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (usuo *UsStateUpdateOne) SetNillableName(s *string) *UsStateUpdateOne {
	if s != nil {
		usuo.SetName(*s)
	}
	return usuo
}

// SetAbbreviation sets the "abbreviation" field.
func (usuo *UsStateUpdateOne) SetAbbreviation(s string) *UsStateUpdateOne {
	usuo.mutation.SetAbbreviation(s)
	return usuo
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (usuo *UsStateUpdateOne) SetNillableAbbreviation(s *string) *UsStateUpdateOne {
	if s != nil {
		usuo.SetAbbreviation(*s)
	}
	return usuo
}

// SetCountryName sets the "country_name" field.
func (usuo *UsStateUpdateOne) SetCountryName(s string) *UsStateUpdateOne {
	usuo.mutation.SetCountryName(s)
	return usuo
}

// SetNillableCountryName sets the "country_name" field if the given value is not nil.
func (usuo *UsStateUpdateOne) SetNillableCountryName(s *string) *UsStateUpdateOne {
	if s != nil {
		usuo.SetCountryName(*s)
	}
	return usuo
}

// SetCountryIso3 sets the "country_iso3" field.
func (usuo *UsStateUpdateOne) SetCountryIso3(s string) *UsStateUpdateOne {
	usuo.mutation.SetCountryIso3(s)
	return usuo
}

// SetNillableCountryIso3 sets the "country_iso3" field if the given value is not nil.
func (usuo *UsStateUpdateOne) SetNillableCountryIso3(s *string) *UsStateUpdateOne {
	if s != nil {
		usuo.SetCountryIso3(*s)
	}
	return usuo
}

// Mutation returns the UsStateMutation object of the builder.
func (usuo *UsStateUpdateOne) Mutation() *UsStateMutation {
	return usuo.mutation
}

// Where appends a list predicates to the UsStateUpdate builder.
func (usuo *UsStateUpdateOne) Where(ps ...predicate.UsState) *UsStateUpdateOne {
	usuo.mutation.Where(ps...)
	return usuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (usuo *UsStateUpdateOne) Select(field string, fields ...string) *UsStateUpdateOne {
	usuo.fields = append([]string{field}, fields...)
	return usuo
}

// Save executes the query and returns the updated UsState entity.
func (usuo *UsStateUpdateOne) Save(ctx context.Context) (*UsState, error) {
	usuo.defaults()
	return withHooks(ctx, usuo.sqlSave, usuo.mutation, usuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (usuo *UsStateUpdateOne) SaveX(ctx context.Context) *UsState {
	node, err := usuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (usuo *UsStateUpdateOne) Exec(ctx context.Context) error {
	_, err := usuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usuo *UsStateUpdateOne) ExecX(ctx context.Context) {
	if err := usuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usuo *UsStateUpdateOne) defaults() {
	if _, ok := usuo.mutation.UpdatedAt(); !ok {
		v := usstate.UpdateDefaultUpdatedAt()
		usuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usuo *UsStateUpdateOne) check() error {
	if v, ok := usuo.mutation.Name(); ok {
		if err := usstate.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "UsState.name": %w`, err)}
		}
	}
	if v, ok := usuo.mutation.Abbreviation(); ok {
		if err := usstate.AbbreviationValidator(v); err != nil {
			return &ValidationError{Name: "abbreviation", err: fmt.Errorf(`ent: validator failed for field "UsState.abbreviation": %w`, err)}
		}
	}
	if v, ok := usuo.mutation.CountryName(); ok {
		if err := usstate.CountryNameValidator(v); err != nil {
			return &ValidationError{Name: "country_name", err: fmt.Errorf(`ent: validator failed for field "UsState.country_name": %w`, err)}
		}
	}
	if v, ok := usuo.mutation.CountryIso3(); ok {
		if err := usstate.CountryIso3Validator(v); err != nil {
			return &ValidationError{Name: "country_iso3", err: fmt.Errorf(`ent: validator failed for field "UsState.country_iso3": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (usuo *UsStateUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UsStateUpdateOne {
	usuo.modifiers = append(usuo.modifiers, modifiers...)
	return usuo
}

func (usuo *UsStateUpdateOne) sqlSave(ctx context.Context) (_node *UsState, err error) {
	if err := usuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(usstate.Table, usstate.Columns, sqlgraph.NewFieldSpec(usstate.FieldID, field.TypeUUID))
	id, ok := usuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UsState.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := usuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usstate.FieldID)
		for _, f := range fields {
			if !usstate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usstate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := usuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := usuo.mutation.UpdatedAt(); ok {
		_spec.SetField(usstate.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := usuo.mutation.Name(); ok {
		_spec.SetField(usstate.FieldName, field.TypeString, value)
	}
	if value, ok := usuo.mutation.Abbreviation(); ok {
		_spec.SetField(usstate.FieldAbbreviation, field.TypeString, value)
	}
	if value, ok := usuo.mutation.CountryName(); ok {
		_spec.SetField(usstate.FieldCountryName, field.TypeString, value)
	}
	if value, ok := usuo.mutation.CountryIso3(); ok {
		_spec.SetField(usstate.FieldCountryIso3, field.TypeString, value)
	}
	_spec.AddModifiers(usuo.modifiers...)
	_node = &UsState{config: usuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, usuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usstate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	usuo.mutation.done = true
	return _node, nil
}
